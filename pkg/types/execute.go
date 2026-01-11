package types

type PlaceOrderParam struct {
	ProductID int

	// Basic
	Amount    float64
	Price     float64
	Direction OrderDirection

	// Flags
	OrderType  OrderType
	ReduceOnly bool

	Options *OrderOptions

	//Optional
	Expiration   *uint64
	ID           *int64
	SpotLeverage *bool
}

// OrderOptions encapsulates all advanced fields that influence the construction of the Order Appendix.
// These are optional settings for complex order types (TWAP, Isolated Margin, etc.).
type OrderOptions struct {
	// Trigger settings (e.g., Stop Loss, Take Profit, TWAP).
	// If nil, defaults to TriggerTypeNone.
	Trigger *TriggerOptions

	// Margin mode configuration.
	// If nil, defaults to Cross Margin mode.
	Isolated *IsolatedMargin

	// Spot leverage toggle.
	// Indicates if borrowing is allowed for spot trades.
	SpotLeverage *bool
}

// TriggerOptions maps to the "Trigger Type" and "Value" fields in the 128-bit Order Appendix.
type TriggerOptions struct {
	Type TriggerType // e.g., TriggerTypePrice, TriggerTypeTwap, TriggerTypeTwapCustom

	// The following fields are valid only when Type is TriggerTypeTwap or TriggerTypeTwapCustom.

	TwapCount    uint32  // Number of execution intervals.
	TwapSlippage float64 // Maximum acceptable slippage (e.g., 0.01 for 1%).
}

// IsolatedMargin maps to the "Isolated" flag and "Value" field in the 128-bit Order Appendix.
type IsolatedMargin struct {
	// Amount of collateral (margin) to transfer to the isolated subaccount for this order.
	Amount float64
}

type OrderBody struct {
	Sender     string `json:"sender"`
	PriceX18   string `json:"priceX18"`
	Amount     string `json:"amount"`
	Expiration string `json:"expiration"`
	Nonce      string `json:"nonce"`
	Appendix   string `json:"appendix"`
}

type PlaceOrderDetails struct {
	ProductID    int       `json:"product_id"`
	Order        OrderBody `json:"order"`
	Signature    string    `json:"signature"`
	ID           *int64    `json:"id,omitempty"`
	SpotLeverage *bool     `json:"spot_leverage,omitempty"`
}

type PlaceOrderPayload struct {
	PlaceOrder PlaceOrderDetails `json:"place_order"`
}

type PlaceOrderData struct {
	Digest string `json:"digest"`
}

type PlaceOrderResponse = ExecuteResponse[PlaceOrderData]

type PlaceOrdersPayload struct {
	PlaceOrders struct {
		Orders        []PlaceOrderDetails `json:"orders"`
		StopOnFailure bool                `json:"stop_on_failure"`
	} `json:"place_orders"`
}
type PlaceOrdersData []struct {
	Digest *string `json:"digest"`
	Error  *string `json:"error"`
}

type PlaceOrdersResponse = ExecuteResponse[PlaceOrdersData]

// ==========================================
// Cancel Orders
// ==========================================

type CancelOrdersBody struct {
	Sender     string   `json:"sender"`
	ProductIds []int    `json:"productIds"`
	Digests    []string `json:"digests"`
	Nonce      string   `json:"nonce"`
}

type CancelOrdersPayload struct {
	CancelOrders struct {
		Tx        CancelOrdersBody `json:"tx"`
		Signature string           `json:"signature"`
	} `json:"cancel_orders"`
}

type CancelOrdersParam struct {
	Params []CancelOrderParam
}
type CancelOrderParam struct {
	ProductID int
	Digest    string
}

type CancelledOrder struct {
	ProductID      int    `json:"product_id"`
	Sender         string `json:"sender"`
	PriceX18       string `json:"price_x18"`
	Amount         string `json:"amount"`
	Expiration     string `json:"expiration"`
	OrderType      string `json:"order_type"`
	Nonce          string `json:"nonce"`
	UnfilledAmount string `json:"unfilled_amount"`
	Digest         string `json:"digest"`
	Appendix       string `json:"appendix"`
	PlacedAt       int64  `json:"placed_at"`
}

type CancelOrdersData struct {
	CancelledOrders []CancelledOrder `json:"cancelled_orders"`
}

type CancelOrdersResponse = ExecuteResponse[CancelOrdersData]

type CancelProductOrdersTx struct {
	Sender     string `json:"sender"`
	ProductIDs []int  `json:"productIds"`
	Nonce      string `json:"nonce"`
}

type CancelProductOrdersPayload struct {
	CancelProductOrders struct {
		Tx        CancelProductOrdersTx `json:"tx"`
		Signature string                `json:"signature"`
	} `json:"cancel_product_orders"`
}

type CancelAndPlaceOrderPayload struct {
	CancelAndPlace struct {
		CancelTx struct {
			Sender     string   `json:"sender"`
			ProductIds []int    `json:"productIds"`
			Digests    []string `json:"digests"`
			Nonce      string   `json:"nonce"`
		} `json:"cancel_tx"`
		CancelSignature string            `json:"cancel_signature"`
		PlaceOrder      PlaceOrderDetails `json:"place_order"`
	} `json:"cancel_and_place"`
}

type WithdrawTx struct {
	Sender    string `json:"sender"`
	ProductID int    `json:"productId"`
	Amount    string `json:"amount"`
	Nonce     string `json:"nonce"`
}

type WithdrawCollateralPayload struct {
	WithdrawCollateral struct {
		Tx           WithdrawTx `json:"tx"`
		Signature    string     `json:"signature"`
		SpotLeverage *bool      `json:"spot_leverage,omitempty"`
	} `json:"withdraw_collateral"`
}

type WithdrawCollateralResponse = ExecuteResponse[interface{}]
