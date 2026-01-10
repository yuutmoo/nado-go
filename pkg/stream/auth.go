package stream

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"nado/pkg/signer"
	"time"
)

func (c *StreamClient) Authenticate(ctx context.Context) error {
	if c.signer == nil {
		return errors.New("signer not initialized")
	}
	expiration := uint64(time.Now().Add(30 * time.Second).UnixMilli())
	senderStr := c.signer.SubAccount()

	var senderBytes32 [32]byte
	senderBytes, _ := hex.DecodeString(senderStr[2:])
	copy(senderBytes32[:], senderBytes)

	authReq := &signer.StreamAuthRequest{
		Sender:            senderBytes32,
		Expiration:        expiration,
		VerifyingContract: c.network.EndPoint,
	}
	sig, err := c.signer.SignTypedData(authReq)
	if err != nil {
		return fmt.Errorf("stream: auth signature failed: %w", err)
	}

	authTx := AuthTx{
		Sender:     senderStr,
		Expiration: fmt.Sprintf("%d", expiration),
	}

	req := AuthRequest{
		Method:    "authenticate",
		ID:        c.requestID.Add(1),
		Signature: sig,
		Tx:        authTx,
	}

	if err := c.writeJSON(ctx, req); err != nil {
		return fmt.Errorf("stream: send auth failed: %w", err)
	}

	fmt.Printf("stream: authentication request sent for %s\n", senderStr)
	return nil
}
