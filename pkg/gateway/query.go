package gateway

import (
	"context"
	"encoding/json"
	"fmt"
	"nado/pkg/common"
	"nado/pkg/types"
	"net/url"
)

// GetSystemStatus Gets status of offchain sequencer.
func (c *GatewayClient) GetSystemStatus(ctx context.Context) (*types.SystemStatusResponse, error) {

	params := url.Values{}
	params.Set("type", "status")

	targetURL := c.buildGatewayURL(common.PathQuery)

	var resp types.SystemStatusResponse

	if err := c.get(ctx, targetURL, params, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil

}

// GetContracts Get information about core Nado contracts.
func (c *GatewayClient) GetContracts(ctx context.Context) (*types.ContractsResponse, error) {
	params := url.Values{}
	params.Set("type", "contracts")

	targetURL := c.buildGatewayURL(common.PathQuery)

	var resp types.ContractsResponse

	if err := c.get(ctx, targetURL, params, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil
}

// GetNonces Get execute nonces for a particular address.
func (c *GatewayClient) GetNonces(ctx context.Context, address string) (*types.NoncesResponse, error) {
	params := url.Values{}
	params.Set("type", "nonces")
	params.Set("address", address)

	targetURL := c.buildGatewayURL(common.PathQuery)

	var resp types.NoncesResponse

	if err := c.get(ctx, targetURL, params, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil
}

// GetOrder retrieves the status and details of a specific order by its digest
func (c *GatewayClient) GetOrder(ctx context.Context, productID int, digest string) (*types.OrderResponse, error) {
	params := url.Values{}
	params.Set("type", "order")
	params.Set("product_id", fmt.Sprintf("%d", productID))
	params.Set("digest", digest)

	targetURL := c.buildGatewayURL(common.PathQuery)

	var resp types.OrderResponse

	if err := c.get(ctx, targetURL, params, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil
}

// GetSubaccountInfo retrieves the detailed state of a subaccount.
//
// Parameters:
//   - subaccount: The hex string address of the subaccount (e.g., "0x...").
//   - simulations: Optional list of transactions to simulate (e.g., ApplyDelta). Pass nil if not needed.
//   - withPreState: If true, the response will include the subaccount state BEFORE the simulated transactions in the 'pre_state' field.
func (c *GatewayClient) GetSubaccountInfo(ctx context.Context, subaccount string, simulations []types.SimulateTx, withPreState bool) (*types.SubaccountInfoResponse, error) {

	reqPayload := types.SubaccountInfoRequest{
		Type:       "subaccount_info",
		Subaccount: subaccount,
		PreState:   "false",
	}

	if withPreState {
		reqPayload.PreState = "true"
	}

	// The API requires the 'txns' field to be a JSON-encoded STRING, not a raw JSON array.
	// e.g., "txns": "[{\"apply_delta\":...}]"
	if len(simulations) > 0 {
		txnsBytes, err := json.Marshal(simulations)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal simulations: %w", err)
		}
		reqPayload.Txns = string(txnsBytes)
	}

	targetURL := c.buildGatewayURL(common.PathQuery)

	var resp types.SubaccountInfoResponse

	if err := c.post(ctx, targetURL, reqPayload, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil
}

// GetIsolatedPositions retrieves isolated margin positions for a subaccount.
// Endpoint: GET /query?type=isolated_positions&subaccount={subaccount}
func (c *GatewayClient) GetIsolatedPositions(ctx context.Context, subaccount string) (*types.IsolatedPositionsResponse, error) {
	params := url.Values{}
	params.Set("type", "isolated_positions")
	params.Set("subaccount", subaccount)

	targetURL := c.buildGatewayURL(common.PathQuery)

	var resp types.IsolatedPositionsResponse

	if err := c.get(ctx, targetURL, params, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil
}

// GetMarketLiquidity retrieves the order book (bids and asks) for a specific product.
// depth: The number of levels to retrieve (e.g., 10, 20).
func (c *GatewayClient) GetMarketLiquidity(ctx context.Context, productID int, depth int) (*types.MarketLiquidityResponse, error) {
	params := url.Values{}
	params.Set("type", "market_liquidity")
	params.Set("product_id", fmt.Sprintf("%d", productID))

	// Only set depth if a valid number is provided
	if depth > 0 {
		params.Set("depth", fmt.Sprintf("%d", depth))
	}

	targetURL := c.buildGatewayURL(common.PathQuery)

	var resp types.MarketLiquidityResponse

	if err := c.get(ctx, targetURL, params, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil
}

// GetSymbols retrieves trading specifications.
// productType: Optional filter ("spot", "perp", or "").
// productIDs: Optional filter list of IDs. IMPORTANT: If provided, this forces a POST request.
func (c *GatewayClient) GetSymbols(ctx context.Context, productType string, productIDs []int) (*types.SymbolsResponse, error) {
	targetURL := c.buildGatewayURL(common.PathQuery)
	var resp types.SymbolsResponse

	if len(productIDs) > 0 {
		reqPayload := types.SymbolsRequest{
			Type:       "symbols",
			ProductIDs: productIDs,
		}
		if productType != "" {
			reqPayload.ProductType = productType
		}

		if err := c.post(ctx, targetURL, reqPayload, &resp); err != nil {
			return nil, err
		}

	} else {
		params := url.Values{}
		params.Set("type", "symbols")

		if productType != "" {
			params.Set("product_type", productType)
		}

		if err := c.get(ctx, targetURL, params, &resp); err != nil {
			return nil, err
		}
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil
}

// GetAllProducts retrieves configuration and state for all available products.
func (c *GatewayClient) GetAllProducts(ctx context.Context) (*types.AllProductsResponse, error) {
	params := url.Values{}
	params.Set("type", "all_products")

	targetURL := c.buildGatewayURL(common.PathQuery)

	var resp types.AllProductsResponse

	if err := c.get(ctx, targetURL, params, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil
}

// GetEdgeAllProducts Gets info about all available products across all chains including: product id,
// oracle price, configuration, state, etc.
func (c *GatewayClient) GetEdgeAllProducts(ctx context.Context) (*types.EdgeAllProductsResponse, error) {
	params := url.Values{}
	params.Set("type", "edge_all_products")

	targetURL := c.buildGatewayURL(common.PathQuery)

	var resp types.EdgeAllProductsResponse

	if err := c.get(ctx, targetURL, params, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil
}

// GetMarketPrice retrieves the current best bid and ask prices for a specific product.
func (c *GatewayClient) GetMarketPrice(ctx context.Context, productID int) (*types.MarketPriceResponse, error) {
	params := url.Values{}
	params.Set("type", "market_price")
	params.Set("product_id", fmt.Sprintf("%d", productID))

	targetURL := c.buildGatewayURL(common.PathQuery)

	var resp types.MarketPriceResponse

	if err := c.get(ctx, targetURL, params, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil
}

// GetMaxOrderSize calculates the maximum possible order size.
// Using GET method as preferred for read-only calculations.
func (c *GatewayClient) GetMaxOrderSize(ctx context.Context, params types.MaxOrderSizeParams) (*types.MaxOrderSizeResponse, error) {
	// 1. Build Query Parameters
	queryParams := url.Values{}
	queryParams.Set("type", "max_order_size")
	queryParams.Set("product_id", fmt.Sprintf("%d", params.ProductID))
	queryParams.Set("sender", params.Sender)
	queryParams.Set("price_x18", params.PriceX18)
	queryParams.Set("direction", params.Direction)

	// Handle optional booleans
	// Convert bool true/false to string "true"/"false"
	if params.SpotLeverage {
		queryParams.Set("spot_leverage", "true")
	} else {
		queryParams.Set("spot_leverage", "false")
	}

	if params.ReduceOnly {
		queryParams.Set("reduce_only", "true")
	} else {
		queryParams.Set("reduce_only", "false")
	}

	if params.Isolated {
		queryParams.Set("isolated", "true")
	} else {
		queryParams.Set("isolated", "false")
	}

	// 2. Build URL
	targetURL := c.buildGatewayURL(common.PathQuery)

	var resp types.MaxOrderSizeResponse

	// 3. Execute GET
	// c.get will append the queryParams to the targetURL automatically
	if err := c.get(ctx, targetURL, queryParams, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil
}

func (c *GatewayClient) GetMaxWithdrawable(ctx context.Context, productID int, sender string, spotLeverage bool) (*types.MaxWithdrawableResponse, error) {
	params := url.Values{}
	params.Set("type", "max_withdrawable")
	params.Set("product_id", fmt.Sprintf("%d", productID))
	params.Set("sender", sender)

	if spotLeverage {
		params.Set("spot_leverage", "true")
	} else {
		params.Set("spot_leverage", "false")
	}

	targetURL := c.buildGatewayURL(common.PathQuery)

	var resp types.MaxWithdrawableResponse

	if err := c.get(ctx, targetURL, params, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil
}

// GetMaxNLPMintable calculates the maximum quote amount that can be used to mint NLP tokens.
// spotLeverage: If true, includes spot leverage in the calculation.
func (c *GatewayClient) GetMaxNLPMintable(ctx context.Context, sender string, spotLeverage bool) (*types.MaxNLPMintableResponse, error) {
	params := url.Values{}
	params.Set("type", "max_nlp_mintable")
	params.Set("sender", sender)

	if spotLeverage {
		params.Set("spot_leverage", "true")
	} else {
		params.Set("spot_leverage", "false")
	}

	targetURL := c.buildGatewayURL(common.PathQuery)

	var resp types.MaxNLPMintableResponse

	if err := c.get(ctx, targetURL, params, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil
}

// GetMaxNLPBurnable calculates the maximum amount of NLP tokens that can be burned (redeemed).
func (c *GatewayClient) GetMaxNLPBurnable(ctx context.Context, sender string) (*types.MaxNLPBurnableResponse, error) {
	params := url.Values{}
	params.Set("type", "max_nlp_burnable")
	params.Set("sender", sender)

	targetURL := c.buildGatewayURL(common.PathQuery)

	var resp types.MaxNLPBurnableResponse

	if err := c.get(ctx, targetURL, params, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil
}

// GetNLPPoolInfo retrieves information about Nado Liquidity Provider (NLP) pools.
func (c *GatewayClient) GetNLPPoolInfo(ctx context.Context) (*types.NLPPoolInfoResponse, error) {
	params := url.Values{}
	params.Set("type", "nlp_pool_info")

	targetURL := c.buildGatewayURL(common.PathQuery)

	var resp types.NLPPoolInfoResponse

	if err := c.get(ctx, targetURL, params, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil
}

// GetNLPLockedBalances retrieves locked and unlocked balance details for an NLP subaccount.
func (c *GatewayClient) GetNLPLockedBalances(ctx context.Context, subaccount string) (*types.NLPLockedBalancesResponse, error) {
	params := url.Values{}
	params.Set("type", "nlp_locked_balances")
	params.Set("subaccount", subaccount)

	targetURL := c.buildGatewayURL(common.PathQuery)

	var resp types.NLPLockedBalancesResponse

	if err := c.get(ctx, targetURL, params, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil
}

// GetFeeRates retrieves the trading fee rates and sequencer fees for a specific sender.
func (c *GatewayClient) GetFeeRates(ctx context.Context, sender string) (*types.FeeRatesResponse, error) {
	params := url.Values{}
	params.Set("type", "fee_rates")
	params.Set("sender", sender)

	targetURL := c.buildGatewayURL(common.PathQuery)

	var resp types.FeeRatesResponse

	if err := c.get(ctx, targetURL, params, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil
}

// GetHealthGroups retrieves the configuration of health groups.
// Note: a health group is a perp and spot product whose health is calculated together (e.g. BTC and BTC-PERP).
func (c *GatewayClient) GetHealthGroups(ctx context.Context) (*types.HealthGroupsResponse, error) {
	params := url.Values{}
	params.Set("type", "health_groups")

	targetURL := c.buildGatewayURL(common.PathQuery)

	var resp types.HealthGroupsResponse

	if err := c.get(ctx, targetURL, params, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil
}

// GetLinkedSigner retrieves the linked signer address for a specific subaccount.
// This is useful to verify if a session key or delegate is currently authorized.
func (c *GatewayClient) GetLinkedSigner(ctx context.Context, subaccount string) (*types.LinkedSignerResponse, error) {
	params := url.Values{}
	params.Set("type", "linked_signer")
	params.Set("subaccount", subaccount)

	targetURL := c.buildGatewayURL(common.PathQuery)

	var resp types.LinkedSignerResponse

	if err := c.get(ctx, targetURL, params, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil
}

// GetInsurance retrieves the current balance of the insurance fund.
func (c *GatewayClient) GetInsurance(ctx context.Context) (*types.InsuranceResponse, error) {
	params := url.Values{}
	params.Set("type", "insurance")

	targetURL := c.buildGatewayURL(common.PathQuery)

	var resp types.InsuranceResponse

	if err := c.get(ctx, targetURL, params, &resp); err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return &resp, fmt.Errorf("%s", resp.Error())
	}

	return &resp, nil
}
