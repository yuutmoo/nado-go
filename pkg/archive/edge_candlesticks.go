package archive

import (
	"context"
	"fmt"
)

// EdgeCandlestickParams represents the query parameters for cross-chain edge candlesticks.
// These parameters are identical to standard candlesticks.
type EdgeCandlestickParams struct {
	// ProductId: Required. Id of product to fetch candlesticks for.
	ProductId int `json:"product_id"`

	// Granularity: Required. Granularity value in seconds (e.g., 60, 3600).
	Granularity int `json:"granularity"`

	// MaxTime: Optional. Unix epoch in seconds.
	// Return candlesticks with timestamp <= max_time.
	MaxTime *int64 `json:"max_time,omitempty"`

	// Limit: Optional. Max number of candlesticks to return.
	// Defaults to 100. Max possible of 500.
	Limit *int `json:"limit,omitempty"`
}

// edgeCandlesticksReqWrapper wraps the request under the "edge_candlesticks" key.
type edgeCandlesticksReqWrapper struct {
	EdgeCandlesticks EdgeCandlestickParams `json:"edge_candlesticks"`
}

// GetEdgeCandlesticks queries historical candlesticks by product across all chains.
// This endpoint returns data ordered by timestamp descending.
func (c *ArchiveClient) GetEdgeCandlesticks(ctx context.Context, params EdgeCandlestickParams) ([]Candlestick, error) {
	// 1. Validation
	if params.ProductId <= 0 {
		return nil, fmt.Errorf("archive: valid product_id is required for edge candlesticks")
	}
	if params.Granularity <= 0 {
		return nil, fmt.Errorf("archive: granularity is required for edge candlesticks")
	}

	// 2. Wrap Request
	reqPayload := edgeCandlesticksReqWrapper{
		EdgeCandlesticks: params,
	}

	// 3. Send Request
	// The response JSON uses the "candlesticks" key for the result array.
	var respWrapper candlesticksRespWrapper
	if err := c.post(ctx, c.BaseURL, reqPayload, &respWrapper); err != nil {
		return nil, err
	}

	return respWrapper.Candlesticks, nil
}
