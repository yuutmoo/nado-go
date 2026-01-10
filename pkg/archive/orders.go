package archive

import (
	"context"
	"fmt"
)

// MatchedOrdersParams represents the parameters for querying historical matched orders.
// Reference: Archive API - Request Parameters
type MatchedOrdersParams struct {
	Subaccounts  []string `json:"subaccounts,omitempty"`
	ProductIds   []int    `json:"product_ids,omitempty"`
	Digests      []string `json:"digests,omitempty"`
	TriggerTypes []string `json:"trigger_types,omitempty"`
	Idx          *int     `json:"idx,omitempty"`
	MaxTime      int64    `json:"max_time,omitempty"`
	Limit        int      `json:"limit,omitempty"`
	Isolated     *bool    `json:"isolated,omitempty"`
}

// MatchedOrder represents a single historical matched order.
// Reference: Archive API - Response Structure
type MatchedOrder struct {
	Digest                string `json:"digest"`
	Subaccount            string `json:"subaccount"`
	ProductId             int    `json:"product_id"`
	SubmissionIdx         string `json:"submission_idx"`
	LastFillSubmissionIdx string `json:"last_fill_submission_idx"`

	Amount string `json:"amount"`

	PriceX18 string `json:"price_x18"`

	BaseFilled string `json:"base_filled"`

	QuoteFilled string `json:"quote_filled"`

	Fee          string `json:"fee"`
	ClosedAmount string `json:"closed_amount"`
	RealizedPnl  string `json:"realized_pnl"`
	Expiration   string `json:"expiration"`
	Appendix     string `json:"appendix"`
	Nonce        string `json:"nonce"`
	Isolated     bool   `json:"isolated"`
}

// matchedOrdersReqWrapper is an internal struct used to wrap the request payload.
// The API expects the parameters to be nested under an "orders" key.
type matchedOrdersReqWrapper struct {
	Orders MatchedOrdersParams `json:"orders"`
}

// matchedOrdersRespWrapper is an internal struct used to unwrap the response payload.
// The API returns the data nested under an "orders" key.
type matchedOrdersRespWrapper struct {
	Orders []MatchedOrder `json:"orders"`
}

// GetMatchedOrders queries matched orders for subaccounts with various filters.
func (c *ArchiveClient) GetMatchedOrders(ctx context.Context, params MatchedOrdersParams) ([]MatchedOrder, error) {
	hasDigests := len(params.Digests) > 0
	hasSubaccounts := len(params.Subaccounts) > 0

	if !hasDigests && !hasSubaccounts {
		return nil, fmt.Errorf("archive: must provide either 'subaccounts' or 'digests'")
	}

	if hasDigests && hasSubaccounts {
		return nil, fmt.Errorf("archive: cannot provide both 'subaccounts' and 'digests' at the same time")
	}

	if hasDigests {
		if len(params.ProductIds) > 0 {
			return nil, fmt.Errorf("archive: cannot use 'product_ids' when querying by 'digests'")
		}
		if params.MaxTime > 0 {
			return nil, fmt.Errorf("archive: cannot use 'max_time' when querying by 'digests'")
		}

	}

	reqPayload := matchedOrdersReqWrapper{
		Orders: params,
	}

	var respWrapper matchedOrdersRespWrapper

	if err := c.post(ctx, c.BaseURL, reqPayload, &respWrapper); err != nil {
		return nil, err
	}

	return respWrapper.Orders, nil
}
