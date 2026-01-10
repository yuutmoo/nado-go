package stream

type WSStream struct {
	Type        string  `json:"type"`
	ProductID   *int    `json:"product_id,omitempty"`
	Granularity *int    `json:"granularity,omitempty"`
	Subaccount  *string `json:"subaccount,omitempty"`
}

type WSRequest struct {
	Method string   `json:"method"`
	Stream WSStream `json:"stream"`
	ID     int64    `json:"id"`
}

type WSResponse struct {
	Result interface{} `json:"result"`
	Id     int64       `json:"id"`
	Type   string      `json:"type"`
	Error  string      `json:"error"`
}

type WSTradeData struct {
	Type         string `json:"type"`
	Timestamp    string `json:"timestamp"`
	ProductId    int    `json:"product_id"`
	Price        string `json:"price"`
	TakerQty     string `json:"taker_qty"`
	MakerQty     string `json:"maker_qty"`
	IsTakerBuyer bool   `json:"is_taker_buyer"`
}

type WSDepthData struct {
	Type             string     `json:"type"`
	LastMaxTimestamp string     `json:"last_max_timestamp"`
	MinTimestamp     string     `json:"min_timestamp"`
	MaxTimestamp     string     `json:"max_timestamp"`
	ProductId        int        `json:"product_id"`
	Bids             [][]string `json:"bids"`
	Asks             [][]string `json:"asks"`
}
type WSBestBidOffer struct {
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	ProductId int    `json:"product_id"`
	BidPrice  string `json:"bid_price"`
	BidQty    string `json:"bid_qty"`
	AskPrice  string `json:"ask_price"`
	AskQty    string `json:"ask_qty"`
}

type WSFill struct {
	Type          string `json:"type"`
	Timestamp     string `json:"timestamp"`
	ProductId     int    `json:"product_id"`
	Subaccount    string `json:"subaccount"`
	OrderDigest   string `json:"order_digest"`
	Id            int    `json:"id"`
	Appendix      string `json:"appendix"`
	FilledQty     string `json:"filled_qty"`
	RemainingQty  string `json:"remaining_qty"`
	OriginalQty   string `json:"original_qty"`
	Price         string `json:"price"`
	IsTaker       bool   `json:"is_taker"`
	IsBid         bool   `json:"is_bid"`
	Fee           string `json:"fee"`
	SubmissionIdx string `json:"submission_idx"`
}

type WSOrderUpdate struct {
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	ProductId int    `json:"product_id"`
	Digest    string `json:"digest"`
	Id        int    `json:"id"`
	Amount    string `json:"amount"`
	Reason    string `json:"reason"`
}

type WSPositionChange struct {
	Type         string `json:"type"`
	Timestamp    string `json:"timestamp"`
	ProductId    int    `json:"product_id"`
	Subaccount   string `json:"subaccount"`
	Isolated     bool   `json:"isolated"`
	Amount       string `json:"amount"`
	VQuoteAmount string `json:"v_quote_amount"`
	Reason       string `json:"reason"`
}

type WSLatestCandlestick struct {
	Type        string `json:"type"`
	Timestamp   string `json:"timestamp"`
	ProductId   int    `json:"product_id"`
	Granularity int    `json:"granularity"`
	OpenX18     string `json:"open_x18"`
	HighX18     string `json:"high_x18"`
	LowX18      string `json:"low_x18"`
	CloseX18    string `json:"close_x18"`
	Volume      string `json:"volume"`
}

type WSFundingPayment struct {
	Type                      string `json:"type"`
	Timestamp                 int64  `json:"timestamp"`
	ProductId                 int    `json:"product_id"`
	PaymentAmount             string `json:"payment_amount"`
	OpenInterest              string `json:"open_interest"`
	CumulativeFundingLongX18  string `json:"cumulative_funding_long_x18"`
	CumulativeFundingShortX18 string `json:"cumulative_funding_short_x18"`
	Dt                        int    `json:"dt"`
}
type WSFundingRate struct {
	Type           string `json:"type"`
	Timestamp      string `json:"timestamp"`
	ProductId      int    `json:"product_id"`
	FundingRateX18 string `json:"funding_rate_x18"`
	UpdateTime     string `json:"update_time"`
}

type WSLiquidation struct {
	Type       string `json:"type"`
	Timestamp  string `json:"timestamp"`
	ProductIds []int  `json:"product_ids"`
	Liquidator string `json:"liquidator"`
	Liquidatee string `json:"liquidatee"`
	Amount     string `json:"amount"`
	Price      string `json:"price"`
}

type AuthTx struct {
	Sender     string `json:"sender"`
	Expiration string `json:"expiration"`
}

type AuthRequest struct {
	Method    string `json:"method"`
	Tx        AuthTx `json:"tx"`
	Signature string `json:"signature"`
	ID        int64  `json:"id"`
}
