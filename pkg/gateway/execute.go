package gateway

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/jinzhu/copier"
	"github.com/yuutmoo/nado-go/pkg/common"
	"github.com/yuutmoo/nado-go/pkg/signer"
	"github.com/yuutmoo/nado-go/pkg/types"
	"math/big"
)

func (c *GatewayClient) PlaceOrder(ctx context.Context, params *types.PlaceOrderParams) (*types.PlaceOrderData, error) {
	if c.signer == nil {
		return nil, fmt.Errorf("execution failed: no signer configured")
	}

	targetURL := c.buildGatewayURL(common.PathExecute)
	var resp types.PlaceOrderResponse

	po := c.buildPlaceOrder(params)
	payload := &types.PlaceOrderPayload{
		PlaceOrder: *po,
	}

	if err := c.post(ctx, targetURL, payload, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp.Data, fmt.Errorf("%s", resp.Error())
	}

	return &resp.Data, nil
}

func (c *GatewayClient) PlaceOrders(ctx context.Context, orders []*types.PlaceOrderParams, stopOnFailure bool) (*types.PlaceOrdersData, error) {
	if len(orders) == 0 {
		return nil, fmt.Errorf("list is empty")
	}

	placeOrders := make([]types.PlaceOrderDetails, 0, len(orders))
	for _, order := range orders {
		po := c.buildPlaceOrder(order)
		placeOrders = append(placeOrders, *po)
	}
	payload := types.PlaceOrdersPayload{
		PlaceOrders: struct {
			Orders        []types.PlaceOrderDetails `json:"orders"`
			StopOnFailure bool                      `json:"stop_on_failure"`
		}{Orders: placeOrders, StopOnFailure: stopOnFailure},
	}

	targetURL := c.buildGatewayURL(common.PathExecute)

	var resp types.PlaceOrdersResponse
	if err := c.post(ctx, targetURL, payload, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {

		return &resp.Data, fmt.Errorf("%s", resp.Error())
	}
	return &resp.Data, nil
}

func (c *GatewayClient) buildPlaceOrder(params *types.PlaceOrderParams) *types.PlaceOrderDetails {

	senderHex := c.signer.SubAccount()

	priceX18 := common.FloatToX18(params.Price, params.PriceTick)

	finalAmount := params.Amount
	if params.Amount < 0 {
		finalAmount = -params.Amount
	}

	if params.Direction == types.DirectionShort {
		finalAmount = -finalAmount
	}

	amountX18 := common.FloatToX18(finalAmount, params.AmountTick)

	nonce := common.GenerateNonce()

	expiration := uint64(4294967295)
	if params.Expiration != nil {
		expiration = *params.Expiration
	}

	var senderBytes32 [32]byte
	senderBytes, _ := hex.DecodeString(senderHex[2:])
	copy(senderBytes32[:], senderBytes)

	appendix := common.BuildAppendix(params)

	internalOrder := &signer.OrderSignRequest{
		ProductID:  params.ProductID,
		Sender:     senderBytes32,
		PriceX18:   priceX18,
		Amount:     amountX18,
		Expiration: expiration,
		Nonce:      nonce,
		Appendix:   appendix,
	}

	sig, err := c.signer.SignTypedData(internalOrder)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	order := types.OrderBody{
		Sender:     senderHex,
		PriceX18:   priceX18,
		Amount:     amountX18,
		Expiration: fmt.Sprintf("%d", expiration),
		Nonce:      fmt.Sprintf("%d", nonce),
		Appendix:   appendix.String(),
	}

	return &types.PlaceOrderDetails{
		ProductID:    params.ProductID,
		Order:        order,
		Signature:    sig,
		ID:           params.ID,
		SpotLeverage: params.SpotLeverage,
	}

}

func (c *GatewayClient) CancelOrders(ctx context.Context, params *types.CancelOrdersParam) (*types.CancelOrdersData, error) {

	if c.signer == nil {
		return nil, fmt.Errorf("execution failed: no signer configured")
	}
	if len(params.Params) == 0 {
		return nil, fmt.Errorf("param is empty")
	}
	targetURL := c.buildGatewayURL(common.PathExecute)

	var resp types.CancelOrdersResponse
	payload := c.buildCancelOrdersPayload(params)

	if err := c.post(ctx, targetURL, payload, &resp); err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return &resp.Data, fmt.Errorf("%s", resp.Error())

	}
	return &resp.Data, nil

}

func (c *GatewayClient) CancelProductOrders(ctx context.Context, productIDs []int) (*types.CancelOrdersData, error) {
	if len(productIDs) == 0 {
		return nil, fmt.Errorf("productIDs is empty")
	}
	targetURL := c.buildGatewayURL(common.PathExecute)
	payload := c.buildCancelProductOrdersPayload(productIDs)
	var resp types.CancelOrdersResponse

	if err := c.post(ctx, targetURL, payload, &resp); err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return &resp.Data, fmt.Errorf("%s", resp.Error())
	}
	return &resp.Data, nil
}

func (c *GatewayClient) CancelAndPlaceOrder(ctx context.Context, orderParam *types.PlaceOrderParams, cancelOrdersParam *types.CancelOrdersParam) (*types.PlaceOrderData, error) {
	if len(cancelOrdersParam.Params) == 0 {
		return nil, nil
	}
	payload := c.buildCancelAndPlaceOrderPayload(orderParam, cancelOrdersParam)

	var resp types.PlaceOrderResponse

	targetURL := c.buildGatewayURL(common.PathExecute)

	if err := c.post(ctx, targetURL, payload, &resp); err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {

		return &resp.Data, fmt.Errorf("%s", resp.Error())
	}
	return &resp.Data, nil

}

func (c *GatewayClient) buildCancelAndPlaceOrderPayload(orderParam *types.PlaceOrderParams, cancelOrdersParam *types.CancelOrdersParam) *types.CancelAndPlaceOrderPayload {

	cancelPayload := c.buildCancelOrdersPayload(cancelOrdersParam)
	orderPayload := c.buildPlaceOrder(orderParam)

	var cancelAndPlaceOrder types.CancelAndPlaceOrderPayload
	cancelAndPlaceOrder.CancelAndPlace.PlaceOrder = *orderPayload
	copier.Copy(&cancelAndPlaceOrder.CancelAndPlace.CancelTx, &cancelPayload.CancelOrders.Tx)
	cancelAndPlaceOrder.CancelAndPlace.CancelSignature = cancelPayload.CancelOrders.Signature
	return &cancelAndPlaceOrder
}

func (c *GatewayClient) buildCancelProductOrdersPayload(productIDs []int) *types.CancelProductOrdersPayload {

	senderStr := c.signer.SubAccount()
	var senderBytes32 [32]byte
	senderBytes, _ := hex.DecodeString(senderStr[2:])
	copy(senderBytes32[:], senderBytes)

	nonce := common.GenerateNonce()

	pIdsUint32 := make([]*math.HexOrDecimal256, len(productIDs))
	for i, pid := range productIDs {
		pIdsUint32[i] = math.NewHexOrDecimal256(int64(pid))
	}

	CancellationProducts := &signer.CancelProductOrdersSignRequest{
		Sender:            senderBytes32,
		ProductIDs:        pIdsUint32,
		Nonce:             nonce,
		VerifyingContract: c.network.EndPoint,
	}

	signature, err := c.signer.SignTypedData(CancellationProducts)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &types.CancelProductOrdersPayload{
		CancelProductOrders: struct {
			Tx        types.CancelProductOrdersTx `json:"tx"`
			Signature string                      `json:"signature"`
		}{
			Tx: struct {
				Sender     string `json:"sender"`
				ProductIDs []int  `json:"productIds"`
				Nonce      string `json:"nonce"`
			}{Sender: senderStr, ProductIDs: productIDs, Nonce: fmt.Sprintf("%d", nonce)},
			Signature: signature,
		},
	}
}

func (c *GatewayClient) buildCancelOrdersPayload(params *types.CancelOrdersParam) *types.CancelOrdersPayload {
	senderStr := c.signer.SubAccount()
	var senderBytes32 [32]byte
	senderBytes, _ := hex.DecodeString(senderStr[2:])
	copy(senderBytes32[:], senderBytes)

	nonce := common.GenerateNonce()

	pIdsUint32 := make([]*math.HexOrDecimal256, len(params.Params))
	pIds := make([]int, len(params.Params))
	digestsBytes32 := make([][32]byte, len(params.Params))
	digests := make([]string, len(params.Params))

	for i, param := range params.Params {
		pIds[i] = param.ProductID
		pIdsUint32[i] = math.NewHexOrDecimal256(int64(pIds[i]))
		digests[i] = param.Digest
		dBytes, err := hex.DecodeString(param.Digest[2:])
		if err != nil {
			return nil
		}
		copy(digestsBytes32[i][:], dBytes)
	}

	cancellation := &signer.CancelSignRequest{
		Sender:            senderBytes32,
		ProductIDs:        pIdsUint32,
		Digests:           digestsBytes32,
		Nonce:             nonce,
		VerifyingContract: c.network.EndPoint,
	}
	signature, err := c.signer.SignTypedData(cancellation)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &types.CancelOrdersPayload{
		CancelOrders: struct {
			Tx        types.CancelOrdersBody `json:"tx"`
			Signature string                 `json:"signature"`
		}{
			Tx: types.CancelOrdersBody{
				Sender:     senderStr,
				ProductIds: pIds,
				Digests:    digests,
				Nonce:      fmt.Sprintf("%d", nonce),
			},
			Signature: signature,
		},
	}
}

func (c *GatewayClient) WithdrawCollateral(ctx context.Context, productId int, amount float64, tokenDecimals int32, SpotLeverage bool) error {

	senderStr := c.signer.SubAccount()
	var senderBytes32 [32]byte
	senderBytes, _ := hex.DecodeString(senderStr[2:])
	copy(senderBytes32[:], senderBytes)

	nonce := common.GenerateNonce()
	amountStr := common.FloatToIntString(amount, tokenDecimals)
	bInt, ok := new(big.Int).SetString(amountStr, 10)
	if !ok {
		return fmt.Errorf("invalid amount format")
	}
	amountVal := (*math.HexOrDecimal256)(bInt)
	internal := &signer.WithdrawCollateralSignRequest{
		Sender:            senderBytes32,
		ProductID:         math.NewHexOrDecimal256(int64(productId)),
		Nonce:             nonce,
		Amount:            amountVal,
		VerifyingContract: c.network.EndPoint,
	}
	signature, err := c.signer.SignTypedData(internal)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	payload := &types.WithdrawCollateralPayload{
		WithdrawCollateral: struct {
			Tx           types.WithdrawTx `json:"tx"`
			Signature    string           `json:"signature"`
			SpotLeverage *bool            `json:"spot_leverage,omitempty"`
		}{
			Tx: types.WithdrawTx{
				Sender:    senderStr,
				ProductID: productId,
				Amount:    amountStr,
				Nonce:     fmt.Sprintf("%d", nonce),
			},
			Signature:    signature,
			SpotLeverage: &SpotLeverage,
		},
	}
	targetURL := c.buildGatewayURL(common.PathExecute)
	var resp types.WithdrawCollateralResponse
	if err := c.post(ctx, targetURL, payload, &resp); err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return fmt.Errorf("%s", resp.Error())
	}

	return nil

}
