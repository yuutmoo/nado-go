package stream

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/yuutmoo/nado-go/log"
	"github.com/yuutmoo/nado-go/pkg/signer"
	"time"
)

func (c *StreamClient) Authenticate(ctx context.Context) error {
	if c.Signer == nil {
		return errors.New("signer not initialized")
	}
	expiration := uint64(time.Now().Add(30 * time.Second).UnixMilli())
	senderStr := c.Signer.SubAccount()

	var senderBytes32 [32]byte
	senderBytes, _ := hex.DecodeString(senderStr[2:])
	copy(senderBytes32[:], senderBytes)

	authReq := &signer.StreamAuthRequest{
		Sender:            senderBytes32,
		Expiration:        expiration,
		VerifyingContract: c.network.EndPoint,
	}
	sig, err := c.Signer.SignTypedData(authReq)
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

	log.Printf("stream: authentication request sent for %s\n", senderStr)
	return nil
}
