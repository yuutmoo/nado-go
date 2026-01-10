package signer

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"math/big"
)

type TypedDataRequest interface {
	GetPrimaryType() string
	GetTypes() apitypes.Types
	GetMessage() map[string]interface{}
	GetVerifyingContract() common.Address
}

func (s *Signer) SignTypedData(req TypedDataRequest) (string, error) {
	domain := apitypes.TypedDataDomain{
		Name:              "Nado",
		Version:           "0.0.1",
		ChainId:           math.NewHexOrDecimal256(s.chainID),
		VerifyingContract: req.GetVerifyingContract().Hex(),
	}

	typedData := apitypes.TypedData{
		Types:       req.GetTypes(),
		PrimaryType: req.GetPrimaryType(),
		Domain:      domain,
		Message:     req.GetMessage(),
	}

	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return "", fmt.Errorf("hash domain failed: %w", err)
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return "", fmt.Errorf("hash message failed: %w", err)
	}

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hash := crypto.Keccak256(rawData)

	signatureBytes, err := crypto.Sign(hash, s.privateKey)
	if err != nil {
		return "", fmt.Errorf("sign failed: %w", err)
	}

	if signatureBytes[64] < 27 {
		signatureBytes[64] += 27
	}

	return "0x" + hex.EncodeToString(signatureBytes), nil
}

type OrderSignRequest struct {
	Sender     [32]byte
	ProductID  int
	PriceX18   string
	Amount     string
	Expiration uint64
	Nonce      uint64
	Appendix   *big.Int
}

func (r *OrderSignRequest) GetPrimaryType() string {
	return "Order"
}

func (r *OrderSignRequest) GetTypes() apitypes.Types {
	return apitypes.Types{
		"EIP712Domain": {
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
		"Order": {
			{Name: "sender", Type: "bytes32"},
			{Name: "priceX18", Type: "int128"},
			{Name: "amount", Type: "int128"},
			{Name: "expiration", Type: "uint64"},
			{Name: "nonce", Type: "uint64"},
			{Name: "appendix", Type: "uint128"},
		},
	}
}
func (r *OrderSignRequest) GetMessage() map[string]interface{} {
	return map[string]interface{}{
		"sender":     r.Sender,
		"priceX18":   r.PriceX18,
		"amount":     r.Amount,
		"expiration": math.NewHexOrDecimal256(int64(r.Expiration)),
		"nonce":      math.NewHexOrDecimal256(int64(r.Nonce)),
		"appendix":   r.Appendix,
	}
}

func (r *OrderSignRequest) GetVerifyingContract() common.Address {
	return common.BytesToAddress(big.NewInt(int64(r.ProductID)).Bytes())
}

type CancelSignRequest struct {
	Sender            [32]byte
	ProductIDs        []*math.HexOrDecimal256
	Digests           [][32]byte
	VerifyingContract string
	Nonce             uint64
}

func (r *CancelSignRequest) GetPrimaryType() string { return "Cancellation" }

func (r *CancelSignRequest) GetTypes() apitypes.Types {
	return apitypes.Types{
		"EIP712Domain": {
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
		"Cancellation": {
			{Name: "sender", Type: "bytes32"},
			{Name: "productIds", Type: "uint32[]"},
			{Name: "digests", Type: "bytes32[]"},
			{Name: "nonce", Type: "uint64"},
		},
	}
}

func (r *CancelSignRequest) GetMessage() map[string]interface{} {
	return map[string]interface{}{
		"sender":     r.Sender,
		"productIds": r.ProductIDs,
		"digests":    r.Digests,
		"nonce":      math.NewHexOrDecimal256(int64(r.Nonce)),
	}
}

func (r *CancelSignRequest) GetVerifyingContract() common.Address {
	return common.HexToAddress(r.VerifyingContract)
}

type CancelProductOrdersSignRequest struct {
	Sender            [32]byte
	ProductIDs        []*math.HexOrDecimal256
	Nonce             uint64
	VerifyingContract string
}

func (r *CancelProductOrdersSignRequest) GetPrimaryType() string {
	return "CancellationProducts"
}

func (r *CancelProductOrdersSignRequest) GetTypes() apitypes.Types {
	return apitypes.Types{
		"EIP712Domain": {
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
		"CancellationProducts": {
			{Name: "sender", Type: "bytes32"},
			{Name: "productIds", Type: "uint32[]"},
			{Name: "nonce", Type: "uint64"},
		},
	}
}

func (r *CancelProductOrdersSignRequest) GetMessage() map[string]interface{} {
	return map[string]interface{}{
		"sender":     r.Sender,
		"productIds": r.ProductIDs,
		"nonce":      math.NewHexOrDecimal256(int64(r.Nonce)),
	}
}
func (r *CancelProductOrdersSignRequest) GetVerifyingContract() common.Address {
	return common.HexToAddress(r.VerifyingContract)
}

type WithdrawCollateralSignRequest struct {
	Sender            [32]byte
	ProductID         *math.HexOrDecimal256
	Nonce             uint64
	Amount            *math.HexOrDecimal256
	VerifyingContract string
}

func (r WithdrawCollateralSignRequest) GetPrimaryType() string {
	return "WithdrawCollateral"
}

func (r WithdrawCollateralSignRequest) GetTypes() apitypes.Types {
	return apitypes.Types{
		"EIP712Domain": {
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
		"WithdrawCollateral": {
			{Name: "sender", Type: "bytes32"},
			{Name: "productId", Type: "uint32"},
			{Name: "amount", Type: "uint128"},
			{Name: "nonce", Type: "uint64"},
		},
	}
}

func (r WithdrawCollateralSignRequest) GetMessage() map[string]interface{} {
	return map[string]interface{}{
		"sender":    r.Sender,
		"productId": r.ProductID,
		"nonce":     math.NewHexOrDecimal256(int64(r.Nonce)),
		"amount":    r.Amount,
	}
}

func (r WithdrawCollateralSignRequest) GetVerifyingContract() common.Address {
	return common.HexToAddress(r.VerifyingContract)
}

type StreamAuthRequest struct {
	Sender            [32]byte
	Expiration        uint64
	VerifyingContract string
}

func (r *StreamAuthRequest) GetPrimaryType() string { return "StreamAuthentication" }

func (r *StreamAuthRequest) GetTypes() apitypes.Types {
	return apitypes.Types{
		"EIP712Domain": {
			{Name: "name", Type: "string"},
			{Name: "version", Type: "string"},
			{Name: "chainId", Type: "uint256"},
			{Name: "verifyingContract", Type: "address"},
		},
		"StreamAuthentication": {
			{Name: "sender", Type: "bytes32"},
			{Name: "expiration", Type: "uint64"},
		},
	}
}

func (r *StreamAuthRequest) GetMessage() map[string]interface{} {
	return map[string]interface{}{
		"sender":     r.Sender,
		"expiration": math.NewHexOrDecimal256(int64(r.Expiration)),
	}
}

func (r *StreamAuthRequest) GetVerifyingContract() common.Address {
	return common.HexToAddress(r.VerifyingContract)
}
