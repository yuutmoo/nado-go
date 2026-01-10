package signer

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

type Signer struct {
	privateKey     *ecdsa.PrivateKey
	Address        common.Address
	SubAccountName string
	chainID        int64
}

func NewSigner(privateKeyHex string, subAccountName string, chainID int64) *Signer {
	if strings.HasPrefix(privateKeyHex, "0x") {
		privateKeyHex = privateKeyHex[2:]
	}

	if subAccountName == "" {
		subAccountName = "default"
	}
	pk, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		panic(fmt.Sprintf("failed to parse private key: %v", err))
	}
	return &Signer{
		privateKey:     pk,
		SubAccountName: subAccountName,
		Address:        crypto.PubkeyToAddress(pk.PublicKey),
		chainID:        chainID,
	}

}

func (s *Signer) SubAccount() string {
	return BuildSender(s.Address, s.SubAccountName)
}

// BuildSender constructs the 32-byte sender/subaccount ID
// Format: Address (20 bytes) + SubaccountName (12 bytes, right padded with nulls)
func BuildSender(address common.Address, subaccountName string) string {
	var subNameBytes [12]byte
	copy(subNameBytes[:], subaccountName)

	senderBytes := append(address.Bytes(), subNameBytes[:]...)
	return "0x" + hex.EncodeToString(senderBytes)
}
