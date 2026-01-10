package archive

import (
	"context"
	"fmt"
)

// CandlestickParams represents the query parameters for historical candlesticks.
type CandlestickParams struct {
	// ProductId: Required. ID of product to fetch candlesticks for.
	ProductId int `json:"product_id"`

	// Granularity: Required. Granularity value in seconds.
	// Supported values: 60, 300, 900, 3600, 7200, 14400, 86400, 604800, 2419200.
	Granularity int `json:"granularity"`

	// MaxTime: Optional. Unix epoch in seconds.
	// Only return candlesticks with timestamp <= max_time.
	MaxTime *int64 `json:"max_time,omitempty"`

	// Limit: Optional. Max number of candlesticks to return.
	// Defaults to 100. Max possible value is 500.
	Limit *int `json:"limit,omitempty"`
}

// Candlestick represents a single OHLCV data point.
type Candlestick struct {
	ProductId     int    `json:"product_id"`
	Granularity   int    `json:"granularity"`
	SubmissionIdx string `json:"submission_idx"`
	Timestamp     string `json:"timestamp"` // Unix epoch in seconds
	OpenX18       string `json:"open_x18"`
	HighX18       string `json:"high_x18"`
	LowX18        string `json:"low_x18"`
	CloseX18      string `json:"close_x18"`
	Volume        string `json:"volume"`
}

// candlesticksReqWrapper wraps the request under the "candlesticks" key.
type candlesticksReqWrapper struct {
	Candlesticks CandlestickParams `json:"candlesticks"`
}

// candlesticksRespWrapper unwraps the response from the "candlesticks" key.
type candlesticksRespWrapper struct {
	Candlesticks []Candlestick `json:"candlesticks"`
}

// GetCandlesticks queries historical product candlesticks ordered by timestamp descending.
func (c *ArchiveClient) GetCandlesticks(ctx context.Context, params CandlestickParams) ([]Candlestick, error) {
	// 1. Basic Validation
	if params.ProductId <= 0 {
		return nil, fmt.Errorf("archive: valid product_id is required")
	}
	if params.Granularity <= 0 {
		return nil, fmt.Errorf("archive: granularity is required")
	}

	// 2. Wrap Request
	reqPayload := candlesticksReqWrapper{
		Candlesticks: params,
	}

	// 3. Send Request
	var respWrapper candlesticksRespWrapper
	if err := c.post(ctx, c.BaseURL, reqPayload, &respWrapper); err != nil {
		return nil, err
	}

	return respWrapper.Candlesticks, nil
}
