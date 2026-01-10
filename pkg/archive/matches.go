package archive

import (
	"context"
)

type MatchesParams struct {
	Subaccounts []string `json:"subaccounts,omitempty"`

	ProductIds []int `json:"product_ids,omitempty"`
	Idx        *int  `json:"idx,omitempty"`
	MaxTime    int64 `json:"max_time,omitempty"`

	Limit    *int  `json:"limit,omitempty"`
	Isolated *bool `json:"isolated,omitempty"`
}

type matchesReqWrapper struct {
	Matches MatchesParams `json:"matches"`
}

// Match represents the specific execution details of a trade.
type Match struct {
	Digest string     `json:"digest"`
	Order  MatchOrder `json:"order"`

	BaseFilled            string `json:"base_filled"`
	QuoteFilled           string `json:"quote_filled"`
	Fee                   string `json:"fee"`
	SequencerFee          string `json:"sequencer_fee"`
	CumulativeFee         string `json:"cumulative_fee"`
	CumulativeBaseFilled  string `json:"cumulative_base_filled"`
	CumulativeQuoteFilled string `json:"cumulative_quote_filled"`

	ClosedAmount string `json:"closed_amount,omitempty"`
	RealizedPnl  string `json:"realized_pnl,omitempty"`

	SubmissionIdx string `json:"submission_idx"`
	Isolated      bool   `json:"isolated"`
	IsTaker       bool   `json:"is_taker"`

	PreBalance  MatchBalanceState `json:"pre_balance"`
	PostBalance MatchBalanceState `json:"post_balance"`
}

// MatchOrder represents the order details as they appear in the Match structure.
// Note: JSON tags here differ slightly from the Tx structure (e.g., priceX18 vs price_x18).
type MatchOrder struct {
	Sender     string `json:"sender"`
	PriceX18   string `json:"priceX18"` // Note: API uses camelCase here based on schema
	Amount     string `json:"amount"`
	Expiration string `json:"expiration"`
	Nonce      string `json:"nonce"`
	Appendix   string `json:"appendix,omitempty"`
}

// MatchBalanceState captures the account state before/after the match.
type MatchBalanceState struct {
	Base  MatchBaseBalance `json:"base"`
	Quote interface{}      `json:"quote"` // Usually null or specific quote structure
}

type MatchBaseBalance struct {
	Perp MatchPerpBalance `json:"perp"`
}

type MatchPerpBalance struct {
	ProductId int             `json:"product_id"`
	Balance   MatchRawBalance `json:"balance"`
}

type MatchRawBalance struct {
	Amount                   string `json:"amount"`
	VQuoteBalance            string `json:"v_quote_balance"`
	LastCumulativeFundingX18 string `json:"last_cumulative_funding_x18"`
}

// -----------------------------------------------------------------------------
// Section 2: Transaction Metadata Structures (L1 Data)
// -----------------------------------------------------------------------------

// MatchTx represents the L1 transaction data associated with the match.
// Use SubmissionIdx to correlate this with a Match.
type MatchTx struct {
	Tx            TxDetails `json:"tx"`
	SubmissionIdx string    `json:"submission_idx"`
	Timestamp     string    `json:"timestamp"` // Important: Timestamp is located here
}

type TxDetails struct {
	MatchOrders TxMatchOrders `json:"match_orders"`
}

type TxMatchOrders struct {
	ProductId int          `json:"product_id"`
	Amm       bool         `json:"amm"`
	Taker     TxOrderEntry `json:"taker"`
	Maker     TxOrderEntry `json:"maker"`
}

type TxOrderEntry struct {
	Order     TxOrder `json:"order"`
	Signature string  `json:"signature"`
}

// TxOrder represents the order details as they appear in the Tx structure.
// Note: Types (int64) and Tags (snake_case) differ from MatchOrder.
type TxOrder struct {
	Sender     string `json:"sender"`
	PriceX18   string `json:"price_x18"` // Note: API uses snake_case here
	Amount     string `json:"amount"`
	Expiration int64  `json:"expiration"` // Note: int64 here, string in MatchOrder
	Appendix   string `json:"appendix"`
	Nonce      int64  `json:"nonce"` // Note: int64 here, string in MatchOrder
}

type MatchesResponse struct {
	Matches []Match   `json:"matches"`
	Txs     []MatchTx `json:"txs"`
}

// GetMatches queries historical matches and their associated transactions.
func (c *ArchiveClient) GetMatches(ctx context.Context, params MatchesParams) (*MatchesResponse, error) {

	reqPayload := matchesReqWrapper{
		Matches: params,
	}

	var resp MatchesResponse
	if err := c.post(ctx, c.BaseURL, reqPayload, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
