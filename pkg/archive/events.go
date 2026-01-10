package archive

import (
	"context"
)

// EventsParams represents the filter criteria for querying historical events.
type EventsParams struct {
	// Subaccounts: Optional. Array of bytes32 hex strings.
	Subaccounts []string `json:"subaccounts,omitempty"`

	// ProductIds: Optional. Filter by specific product IDs.
	ProductIds []int `json:"product_ids,omitempty"`

	// EventTypes: Optional. Filter by event type (e.g., "match_orders", "deposit_collateral").
	EventTypes []string `json:"event_types,omitempty"`

	// Idx: Optional. Return events with submission_idx <= idx.
	Idx *int `json:"idx,omitempty"`

	// MaxTime: Optional. Unix epoch in seconds.
	// Used only if Idx is not provided.
	MaxTime int64 `json:"max_time,omitempty"`

	// Limit: Optional. Supports "raw" (max events) or "txs" (max transactions).
	// Example: {"raw": 100}
	Limit map[string]int `json:"limit,omitempty"`

	// Isolated: Optional. true (isolated only), false (cross only), or nil (all).
	Isolated *bool `json:"isolated,omitempty"`
}

// EventsResponse is the root container for the event query results.
type EventsResponse struct {
	Events []Event   `json:"events"`
	Txs    []EventTx `json:"txs"`
}

// -----------------------------------------------------------------------------
// Section 1: Event Data Structures
// -----------------------------------------------------------------------------

type Event struct {
	Subaccount        string `json:"subaccount"`
	ProductId         int    `json:"product_id"`
	SubmissionIdx     string `json:"submission_idx"`
	EventType         string `json:"event_type"`
	Isolated          bool   `json:"isolated"`
	IsolatedProductId *int   `json:"isolated_product_id"`

	PreBalance  EventBalance `json:"pre_balance"`
	PostBalance EventBalance `json:"post_balance"`
	Product     EventProduct `json:"product"`

	NetInterestUnrealized string `json:"net_interest_unrealized"`
	NetInterestCumulative string `json:"net_interest_cumulative"`
	NetFundingUnrealized  string `json:"net_funding_unrealized"`
	NetFundingCumulative  string `json:"net_funding_cumulative"`
	NetEntryUnrealized    string `json:"net_entry_unrealized"`
	NetEntryCumulative    string `json:"net_entry_cumulative"`
	QuoteVolumeCumulative string `json:"quote_volume_cumulative"`
}

type EventBalance struct {
	Spot *EventSpotBalance `json:"spot,omitempty"`
	Perp *EventPerpBalance `json:"perp,omitempty"`
}

type EventSpotBalance struct {
	ProductId int `json:"product_id"`
	Balance   struct {
		Amount                      string `json:"amount"`
		LastCumulativeMultiplierX18 string `json:"last_cumulative_multiplier_x18,omitempty"`
	} `json:"balance"`
}

type EventPerpBalance struct {
	ProductId int `json:"product_id"`
	Balance   struct {
		Amount                   string `json:"amount"`
		VQuoteBalance            string `json:"v_quote_balance"`
		LastCumulativeFundingX18 string `json:"last_cumulative_funding_x18"`
	} `json:"balance"`
}

type EventProduct struct {
	ProductId      int           `json:"product_id"`
	OraclePriceX18 string        `json:"oracle_price_x18"`
	Risk           EventRisk     `json:"risk"`
	Config         EventConfig   `json:"config"`
	State          EventState    `json:"state"`
	BookInfo       EventBookInfo `json:"book_info"`
}

type EventRisk struct {
	LongWeightInitialX18      string `json:"long_weight_initial_x18"`
	ShortWeightInitialX18     string `json:"short_weight_initial_x18"`
	LongWeightMaintenanceX18  string `json:"long_weight_maintenance_x18"`
	ShortWeightMaintenanceX18 string `json:"short_weight_maintenance_x18"`
	PriceX18                  string `json:"price_x18"`
}

type EventConfig struct {
	Token                     string `json:"token"`
	InterestInflectionUtilX18 string `json:"interest_inflection_util_x18"`
	InterestFloorX18          string `json:"interest_floor_x18"`
	InterestSmallCapX18       string `json:"interest_small_cap_x18"`
	InterestLargeCapX18       string `json:"interest_large_cap_x18"`
	WithdrawFeeX18            string `json:"withdraw_fee_x18"`
	MinDepositRateX18         string `json:"min_deposit_rate_x18"`
}

type EventState struct {
	CumulativeDepositsMultiplierX18 string `json:"cumulative_deposits_multiplier_x18"`
	CumulativeBorrowsMultiplierX18  string `json:"cumulative_borrows_multiplier_x18"`
	TotalDepositsNormalized         string `json:"total_deposits_normalized"`
	TotalBorrowsNormalized          string `json:"total_borrows_normalized"`
}

type EventBookInfo struct {
	SizeIncrement     string `json:"size_increment"`
	PriceIncrementX18 string `json:"price_increment_x18"`
	MinSize           string `json:"min_size"`
	CollectedFees     string `json:"collected_fees"`
}

// -----------------------------------------------------------------------------
// Section 2: Transaction Metadata Structures
// -----------------------------------------------------------------------------

type EventTx struct {
	Tx            interface{} `json:"tx"` // Can be match_orders, deposit, etc.
	SubmissionIdx string      `json:"submission_idx"`
	Timestamp     string      `json:"timestamp"`
}

// -----------------------------------------------------------------------------
// Section 3: API Method
// -----------------------------------------------------------------------------

type eventsReqWrapper struct {
	Events EventsParams `json:"events"`
}

// GetEvents queries historical events based on subaccounts, products, and types.
// It returns a combined object of events and their parent transactions.
func (c *ArchiveClient) GetEvents(ctx context.Context, params EventsParams) (*EventsResponse, error) {
	reqPayload := eventsReqWrapper{
		Events: params,
	}

	var resp EventsResponse
	if err := c.post(ctx, c.BaseURL, reqPayload, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
