package archive

import (
	"context"
	"fmt"
)

// FundingRate represents the 24-hour funding information for a perpetual product.
type FundingRate struct {
	ProductId      int    `json:"product_id"`
	FundingRateX18 string `json:"funding_rate_x18"`
	UpdateTime     string `json:"update_time"`
}

// -----------------------------------------------------------------------------
// Section 1: Single Product Request
// -----------------------------------------------------------------------------

type fundingRateReq struct {
	ProductId int `json:"product_id"`
}

type fundingRateReqWrapper struct {
	FundingRate fundingRateReq `json:"funding_rate"`
}

// GetFundingRate queries the 24hr funding rate for a single perpetual product.
// Rate limit: 1200 requests/min (weight = 2).
func (c *ArchiveClient) GetFundingRate(ctx context.Context, productId int) (*FundingRate, error) {
	reqPayload := fundingRateReqWrapper{
		FundingRate: fundingRateReq{ProductId: productId},
	}

	var resp FundingRate
	if err := c.post(ctx, c.BaseURL, reqPayload, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

// -----------------------------------------------------------------------------
// Section 2: Multiple Products Request
// -----------------------------------------------------------------------------

type fundingRatesReq struct {
	ProductIds []int `json:"product_ids"`
}

type fundingRatesReqWrapper struct {
	FundingRates fundingRatesReq `json:"funding_rates"`
}

// GetFundingRates queries the 24hr funding rates for multiple perpetual products.
// The response is returned as a map where the key is the string representation of the product ID.
func (c *ArchiveClient) GetFundingRates(ctx context.Context, productIds []int) (map[string]FundingRate, error) {
	if len(productIds) == 0 {
		return nil, fmt.Errorf("archive: product_ids slice cannot be empty")
	}

	reqPayload := fundingRatesReqWrapper{
		FundingRates: fundingRatesReq{ProductIds: productIds},
	}

	// The API returns a map of product_id -> funding_rate
	var resp map[string]FundingRate
	if err := c.post(ctx, c.BaseURL, reqPayload, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}
