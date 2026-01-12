package gateway

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/jinzhu/copier"
	"github.com/yuutmoo/nado-go/log"
	"github.com/yuutmoo/nado-go/pkg/common"
	"github.com/yuutmoo/nado-go/pkg/signer"
	"github.com/yuutmoo/nado-go/pkg/types"
	"math/big"
)

func (c *GatewayClient) PlaceOrder(ctx context.Context, params *types.PlaceOrderParam) (*types.PlaceOrderData, error) {
	if c.signer == nil {
		return nil, fmt.Errorf("execution failed: no signer configured")
	}

	if err := params.Validate(); err != nil {
		return nil, fmt.Errorf("invalid order parameters: %w", err)
	}

	var resp types.PlaceOrderResponse

	po := c.buildPlaceOrder(params)
	if po == nil {
		return nil, fmt.Errorf("build placeOrder response failed")
	}
	payload := &types.PlaceOrderPayload{
		PlaceOrder: *po,
	}

	targetURL := c.buildGatewayURL(common.PathExecute)
	if params.Trigger != nil {
		targetURL = c.buildTriggerURL(common.PathExecute)
	}

	if err := c.post(ctx, targetURL, payload, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp.Data, fmt.Errorf("%s", resp.Error())
	}

	return &resp.Data, nil
}

func (c *GatewayClient) PlaceOrders(ctx context.Context, orders []*types.PlaceOrderParam, stopOnFailure bool) (*types.PlaceOrdersData, error) {
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
	if orders[0].Trigger != nil {
		targetURL = c.buildTriggerURL(common.PathExecute)
	}

	var resp types.PlaceOrdersResponse
	if err := c.post(ctx, targetURL, payload, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {

		return &resp.Data, fmt.Errorf("%s", resp.Error())
	}
	return &resp.Data, nil
}

func (c *GatewayClient) buildPlaceOrder(params *types.PlaceOrderParam) *types.PlaceOrderDetails {

	var priceTick float64
	var amountTick float64
	if val, ok := c.productCache.Load(params.ProductID); ok {
		info := val.(types.ProductInfo)
		priceTick = info.PriceTick
		amountTick = info.AmountTick
	} else {
		log.Error("product priceTick, amountTick not found in cache")
		return nil
	}

	var triggerObj *types.TriggerCriteria
	if params.Trigger != nil {
		triggerObj = &types.TriggerCriteria{}
		if params.Trigger.PriceTrigger != nil {
			triggerObj.PriceTrigger = params.Trigger.PriceTrigger.ToRequest(amountTick)
		}
		if params.Trigger.TimeTrigger != nil {
			triggerObj.TimeTrigger = params.Trigger.TimeTrigger.ToRequest(amountTick)
		}
	}

	senderHex := c.signer.SubAccount()

	priceX18 := common.FloatToX18(params.Price, priceTick)

	finalAmount := params.Amount
	if params.Amount < 0 {
		finalAmount = -params.Amount
	}

	if params.Direction == types.DirectionShort {
		finalAmount = -finalAmount
	}

	amountX18 := common.FloatToX18(finalAmount, amountTick)

	nonce := common.GenerateNonce()

	expiration := uint64(4294967295)
	if params.Expiration != nil {
		expiration = *params.Expiration
	}

	var senderBytes32 [32]byte
	senderBytes, _ := hex.DecodeString(senderHex[2:])
	copy(senderBytes32[:], senderBytes)

	appendix := BuildAppendix(params)

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
		log.Error(err)
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
		Trigger:      triggerObj,
		Signature:    sig,
		ID:           params.ID,
		SpotLeverage: params.SpotLeverage,
	}

}

func (c *GatewayClient) CancelOrders(ctx context.Context, params *types.CancelOrdersParam, isTrigger bool) (*types.CancelOrdersData, error) {

	if c.signer == nil {
		return nil, fmt.Errorf("execution failed: no signer configured")
	}
	if len(params.Params) == 0 {
		return nil, fmt.Errorf("param is empty")
	}

	var resp types.CancelOrdersResponse
	payload := c.buildCancelOrdersPayload(params)

	targetURL := c.buildGatewayURL(common.PathExecute)
	if isTrigger {
		targetURL = c.buildTriggerURL(common.PathExecute)
	}
	if err := c.post(ctx, targetURL, payload, &resp); err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return &resp.Data, fmt.Errorf("%s", resp.Error())

	}
	return &resp.Data, nil

}

func (c *GatewayClient) CancelProductOrders(ctx context.Context, productIDs []int, isTrigger bool) (*types.CancelOrdersData, error) {
	if len(productIDs) == 0 {
		return nil, fmt.Errorf("productIDs is empty")
	}
	payload := c.buildCancelProductOrdersPayload(productIDs)
	var resp types.CancelOrdersResponse

	targetURL := c.buildGatewayURL(common.PathExecute)
	if isTrigger {
		targetURL = c.buildTriggerURL(common.PathExecute)
	}
	if err := c.post(ctx, targetURL, payload, &resp); err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return &resp.Data, fmt.Errorf("%s", resp.Error())
	}
	return &resp.Data, nil
}

func (c *GatewayClient) CancelAndPlaceOrder(ctx context.Context, orderParam *types.PlaceOrderParam, cancelOrdersParam *types.CancelOrdersParam) (*types.PlaceOrderData, error) {
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

func (c *GatewayClient) buildCancelAndPlaceOrderPayload(orderParam *types.PlaceOrderParam, cancelOrdersParam *types.CancelOrdersParam) *types.CancelAndPlaceOrderPayload {

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
		log.Error(err)
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
		log.Error(err)
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
		log.Error(err)
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

func BuildAppendix(params *types.PlaceOrderParam) *big.Int {

	// Bit Layout (from docs):
	// Version (0-7): 1
	// Isolated (8): 0 (assuming cross margin)
	// Order Type (9-10)
	// Reduce Only (11)
	appendix := big.NewInt(0)

	// 1. Version (0-7): 1
	appendix.Or(appendix, big.NewInt(1))

	// 2. Isolated (8)
	if params.Isolated {
		appendix.SetBit(appendix, 8, 1)
	}

	// 3. Order Type (9-10)
	typeBig := new(big.Int).SetInt64(int64(params.OrderType))
	appendix.Or(appendix, typeBig.Lsh(typeBig, 9))

	// 4. Reduce Only (11)
	if params.ReduceOnly {
		appendix.SetBit(appendix, 11, 1)
	}

	// 5. Trigger Type (12-13)
	triggerType := types.TriggerTypeNone
	if params.Trigger != nil {
		if params.Trigger.TimeTrigger != nil {
			triggerType = types.TriggerTypeTwap
			if len(params.Trigger.TimeTrigger.Amounts) > 0 {
				triggerType = types.TriggerTypeCustomAmount
				appendix.SetBit(appendix, 13, 1)
			}
		} else if params.Trigger.PriceTrigger != nil {
			triggerType = types.TriggerTypePrice
		}
	}

	trigVal := new(big.Int).SetInt64(int64(triggerType))
	appendix.Or(appendix, trigVal.Lsh(trigVal, 12))

	// 6. Value (64-127)
	valuePart := big.NewInt(0)

	if triggerType == types.TriggerTypeTwap || triggerType == types.TriggerTypeCustomAmount {
		slippageX6 := common.ToX6(params.Trigger.TimeTrigger.Slippage)
		valuePart.SetInt64(slippageX6)

		times := len(params.Trigger.TimeTrigger.Amounts)
		if times == 0 {
			times = params.Trigger.TimeTrigger.Times
		}
		timesBig := new(big.Int).SetInt64(int64(times))
		valuePart.Or(valuePart, timesBig.Lsh(timesBig, 32))
	} else if params.Isolated {
		marginX6 := common.ToX6(params.IsolatedMargin)
		valuePart.SetInt64(marginX6)
	}

	if valuePart.Sign() != 0 {
		appendix.Or(appendix, valuePart.Lsh(valuePart, 64))
	}

	return appendix

}
