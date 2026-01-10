package types

// SystemStatusResponse represents the response from GET /query?type=status
type SystemStatusResponse = BaseResponse[string]

// ContractsData contains the protocol configuration details
type ContractsData struct {
	ChainID      string `json:"chain_id"`
	EndpointAddr string `json:"endpoint_addr"`
}

// ContractsResponse represents the response from GET /query?type=contracts
type ContractsResponse = BaseResponse[ContractsData]

// NoncesData contains the current nonce counters for an address
type NoncesData struct {
	TxNonce    string `json:"tx_nonce"`
	OrderNonce string `json:"order_nonce"`
}

// NoncesResponse represents the response from GET /query?type=nonces
type NoncesResponse = BaseResponse[NoncesData]

// OrderData contains the details of a specific order
type OrderData struct {
	ProductID      int    `json:"product_id"`
	Sender         string `json:"sender"`
	PriceX18       string `json:"price_x18"`
	Amount         string `json:"amount"` // Signed amount (positive for buy, negative for sell)
	Expiration     string `json:"expiration"`
	Nonce          string `json:"nonce"`
	UnfilledAmount string `json:"unfilled_amount"`
	Digest         string `json:"digest"`
	PlacedAt       int64  `json:"placed_at"` // Unix timestamp
	Appendix       string `json:"appendix"`
	OrderType      string `json:"order_type"` // e.g., "ioc", "gtc"
}

// OrderResponse represents the response from GET /query?type=order
type OrderResponse = BaseResponse[OrderData]

// ==========================================
// 1. Health & Balance Components
// ==========================================

// HealthDetail represents the assets, liabilities, and health value
type HealthDetail struct {
	Assets      string `json:"assets"`      // int128 string
	Liabilities string `json:"liabilities"` // int128 string
	Health      string `json:"health"`      // int128 string
}

// Balance Details
type SpotBalanceDetail struct {
	Amount string `json:"amount"` // int128 string
}

type PerpBalanceDetail struct {
	Amount                   string `json:"amount"`                      // int128 string (Size)
	VQuoteBalance            string `json:"v_quote_balance"`             // int128 string (Cost/Quote)
	LastCumulativeFundingX18 string `json:"last_cumulative_funding_x18"` // int128 string
}

// Balance Items (Wrapper with Product ID)
type SpotBalanceItem struct {
	ProductID int               `json:"product_id"`
	Balance   SpotBalanceDetail `json:"balance"`
}

type PerpBalanceItem struct {
	ProductID int               `json:"product_id"`
	Balance   PerpBalanceDetail `json:"balance"`
}

// ==========================================
// 2. Product Metadata Components
// ==========================================

// ProductRisk contains risk parameters used for health calculation
type ProductRisk struct {
	LongWeightInitialX18      string `json:"long_weight_initial_x18"`
	ShortWeightInitialX18     string `json:"short_weight_initial_x18"`
	LongWeightMaintenanceX18  string `json:"long_weight_maintenance_x18"`
	ShortWeightMaintenanceX18 string `json:"short_weight_maintenance_x18"`
	PriceX18                  string `json:"price_x18"`
}

// BookInfo contains order book limits
type BookInfo struct {
	SizeIncrement     string `json:"size_increment"`
	PriceIncrementX18 string `json:"price_increment_x18"`
	MinSize           string `json:"min_size"`
	CollectedFees     string `json:"collected_fees"`
}

// Spot Product Specifics
type SpotConfig struct {
	Token                     string `json:"token"` // ERC20 Address
	InterestInflectionUtilX18 string `json:"interest_inflection_util_x18"`
	InterestFloorX18          string `json:"interest_floor_x18"`
	InterestSmallCapX18       string `json:"interest_small_cap_x18"`
	InterestLargeCapX18       string `json:"interest_large_cap_x18"`
	WithdrawFeeX18            string `json:"withdraw_fee_x18"`
	MinDepositRateX18         string `json:"min_deposit_rate_x18"`
}

type SpotState struct {
	CumulativeDepositsMultiplierX18 string `json:"cumulative_deposits_multiplier_x18"`
	CumulativeBorrowsMultiplierX18  string `json:"cumulative_borrows_multiplier_x18"`
	TotalDepositsNormalized         string `json:"total_deposits_normalized"`
	TotalBorrowsNormalized          string `json:"total_borrows_normalized"`
}

type SpotProduct struct {
	ProductID      int         `json:"product_id"`
	OraclePriceX18 string      `json:"oracle_price_x18"`
	Risk           ProductRisk `json:"risk"`
	Config         SpotConfig  `json:"config"`
	State          SpotState   `json:"state"`
	BookInfo       BookInfo    `json:"book_info"`
}

// Perp Product Specifics
type PerpState struct {
	CumulativeFundingLongX18  string `json:"cumulative_funding_long_x18"`
	CumulativeFundingShortX18 string `json:"cumulative_funding_short_x18"`
	AvailableSettle           string `json:"available_settle"`
	OpenInterest              string `json:"open_interest"`
}

type PerpProduct struct {
	ProductID      int         `json:"product_id"`
	OraclePriceX18 string      `json:"oracle_price_x18"`
	Risk           ProductRisk `json:"risk"`
	// Perps might not have 'config' in this view, based on your JSON
	State    PerpState `json:"state"`
	BookInfo BookInfo  `json:"book_info"`
}

// ==========================================
// 3. Main Subaccount Data
// ==========================================

// SubaccountInfoData represents the full state of a subaccount
type SubaccountInfoData struct {
	Subaccount string `json:"subaccount"`
	Exists     bool   `json:"exists"`

	// Healths: [0]=Initial, [1]=Maintenance, [2]=Unweighted
	Healths []HealthDetail `json:"healths"`

	// HealthContributions: A list of lists of strings.
	// The outer index corresponds to the Product ID.
	// The inner list contains 3 values corresponding to the 3 Health types.
	// e.g. HealthContributions[0][1] is Product 0's contribution to Maintenance Health.
	HealthContributions [][]string `json:"health_contributions"`

	SpotCount int `json:"spot_count"`
	PerpCount int `json:"perp_count"`

	SpotBalances []SpotBalanceItem `json:"spot_balances"`
	PerpBalances []PerpBalanceItem `json:"perp_balances"`

	SpotProducts []SpotProduct `json:"spot_products"`
	PerpProducts []PerpProduct `json:"perp_products"`

	// PreState: Recursive field.
	// Populated only when request param pre_state="true" AND txns are provided.
	PreState *SubaccountInfoData `json:"pre_state,omitempty"`
}

// SubaccountInfoResponse matches the standard response wrapper
type SubaccountInfoResponse = BaseResponse[SubaccountInfoData]

// ==========================================
// 4. Simulation Request Types
// ==========================================

// ApplyDeltaParams for "apply_delta" simulation
type ApplyDeltaParams struct {
	ProductID   int    `json:"product_id"`
	Subaccount  string `json:"subaccount"`
	AmountDelta string `json:"amount_delta"`
	VQuoteDelta string `json:"v_quote_delta"`
}

// SimulateTx represents a single simulation transaction
type SimulateTx struct {
	ApplyDelta *ApplyDeltaParams `json:"apply_delta,omitempty"`
}

// subaccountInfoRequest (Internal Payload)
type SubaccountInfoRequest struct {
	Type       string `json:"type"`
	Subaccount string `json:"subaccount"`
	Txns       string `json:"txns,omitempty"`      // JSON stringified list of SimulateTx
	PreState   string `json:"pre_state,omitempty"` // "true" or "false"
}

// ==========================================
// Isolated Positions Types
// ==========================================

// IsolatedPositionData represents a single isolated margin position entry
type IsolatedPositionData struct {
	Subaccount string `json:"subaccount"`

	QuoteBalance SpotBalanceItem `json:"quote_balance"`
	BaseBalance  PerpBalanceItem `json:"base_balance"`

	QuoteProduct SpotProduct `json:"quote_product"`
	BaseProduct  PerpProduct `json:"base_product"`

	// Healths specific to this isolated position
	// Arrays of strings (int128): [0]=Initial, [1]=Maintenance, [2]=Unweighted
	QuoteHealths []string       `json:"quote_healths"`
	BaseHealths  []string       `json:"base_healths"`
	Healths      []HealthDetail `json:"healths"`
}

// IsolatedPositionsDataWrapper wraps the list
type IsolatedPositionsDataWrapper struct {
	IsolatedPositions []IsolatedPositionData `json:"isolated_positions"`
}

// IsolatedPositionsResponse matches the standard response wrapper
type IsolatedPositionsResponse = BaseResponse[IsolatedPositionsDataWrapper]

// MarketLiquidityData represents the order book depth
type MarketLiquidityData struct {
	// Bids and Asks are lists of price levels.
	// Each level is an array of 2 strings: [0]=PriceX18, [1]=AmountX18
	Bids      [][]string `json:"bids"`
	Asks      [][]string `json:"asks"`
	Timestamp string     `json:"timestamp"` // Nanoseconds timestamp string
	ProductID int        `json:"product_id"`
}

// MarketLiquidityResponse represents the response from GET /query?type=market_liquidity
type MarketLiquidityResponse = BaseResponse[MarketLiquidityData]

// SymbolInfo contains trading rules and details for a specific product
type SymbolInfo struct {
	Type                     string `json:"type"` // "spot" or "perp"
	ProductID                int    `json:"product_id"`
	Symbol                   string `json:"symbol"`
	PriceIncrementX18        string `json:"price_increment_x18"`
	SizeIncrement            string `json:"size_increment"`
	MinSize                  string `json:"min_size"`
	MakerFeeRateX18          string `json:"maker_fee_rate_x18"`
	TakerFeeRateX18          string `json:"taker_fee_rate_x18"`
	LongWeightInitialX18     string `json:"long_weight_initial_x18"`
	LongWeightMaintenanceX18 string `json:"long_weight_maintenance_x18"`

	MaxOpenInterestX18 *string `json:"max_open_interest_x18"`
}

type SymbolsRequest struct {
	Type        string `json:"type"`
	ProductType string `json:"product_type,omitempty"` // "spot" | "perp"
	ProductIDs  []int  `json:"product_ids,omitempty"`  // Only available for POST
}

// SymbolsData wraps the map of symbols
type SymbolsData struct {
	// Key is the symbol name (e.g. "WBTC", "BTC-PERP")
	Symbols map[string]SymbolInfo `json:"symbols"`
}

// SymbolsResponse represents the response from GET /query?type=symbols
type SymbolsResponse = BaseResponse[SymbolsData]

// AllProductsData contains lists of all available spot and perp products
type AllProductsData struct {
	SpotProducts []SpotProduct `json:"spot_products"`
	PerpProducts []PerpProduct `json:"perp_products"`
}

// AllProductsResponse represents the response from GET /query?type=all_products
type AllProductsResponse = BaseResponse[AllProductsData]

// EdgeAllProductsData contains product configurations indexed by Chain ID.
type EdgeAllProductsData struct {
	// Key is the Chain ID string (e.g., "763373")
	// Value reuses the AllProductsData structure defined previously
	EdgeAllProducts map[string]AllProductsData `json:"edge_all_products"`
}

// EdgeAllProductsResponse represents the response from GET /query?type=edge_all_products
type EdgeAllProductsResponse = BaseResponse[EdgeAllProductsData]

// MarketPriceData represents the best bid and ask prices for a product
type MarketPriceData struct {
	ProductID int    `json:"product_id"`
	BidX18    string `json:"bid_x18"` // Best bid price
	AskX18    string `json:"ask_x18"` // Best ask price
}

// MarketPriceResponse represents the response from GET /query?type=market_price
type MarketPriceResponse = BaseResponse[MarketPriceData]

// MaxOrderSizeData contains the calculation result
type MaxOrderSizeData struct {
	MaxOrderSize string `json:"max_order_size"` // int128 string
}

// MaxOrderSizeResponse represents the response from query (type=max_order_size)
type MaxOrderSizeResponse = BaseResponse[MaxOrderSizeData]

// MaxOrderSizeParams is the input struct for the client method.
type MaxOrderSizeParams struct {
	ProductID    int
	Sender       string // Subaccount address
	PriceX18     string
	Direction    string // "long" or "short"
	SpotLeverage bool   // Default false
	ReduceOnly   bool   // Default false
	Isolated     bool   // Default false
}

// MaxWithdrawableData contains the calculation result
type MaxWithdrawableData struct {
	MaxWithdrawable string `json:"max_withdrawable"` // int128 string
}

// MaxWithdrawableResponse represents the response from GET /query?type=max_withdrawable
type MaxWithdrawableResponse = BaseResponse[MaxWithdrawableData]

// MaxNLPMintableData contains the max quote amount calculation
type MaxNLPMintableData struct {
	MaxQuoteAmount string `json:"max_quote_amount"` // int128 string
}

// MaxNLPMintableResponse represents the response from GET /query?type=max_nlp_mintable
type MaxNLPMintableResponse = BaseResponse[MaxNLPMintableData]

// MaxNLPBurnableData contains the max NLP amount calculation
type MaxNLPBurnableData struct {
	MaxNLPAmount string `json:"max_nlp_amount"` // int128 string
}

// MaxNLPBurnableResponse represents the response from GET /query?type=max_nlp_burnable
type MaxNLPBurnableResponse = BaseResponse[MaxNLPBurnableData]

// ==========================================
// NLP Pool Types
// ==========================================

// NLPPoolHealth represents the health metrics specific to the NLP pool view
type NLPPoolHealth struct {
	Assets            string `json:"assets"`
	Liabilities       string `json:"liabilities"`
	InitialHealth     string `json:"initial_health"`
	MaintenanceHealth string `json:"maintenance_health"`
}

// NLPPoolSubaccountInfo is a simplified subaccount view inside the NLP pool response.
// Note: Structurally different from the main SubaccountInfoData (health is an object here).
type NLPPoolSubaccountInfo struct {
	Subaccount   string            `json:"subaccount"`
	Exists       bool              `json:"exists"`
	Health       NLPPoolHealth     `json:"health"`
	SpotBalances []SpotBalanceItem `json:"spot_balances"` // Reuse existing type
	PerpBalances []PerpBalanceItem `json:"perp_balances"` // Reuse existing type
}

// NLPPool represents a single liquidity pool
type NLPPool struct {
	PoolID           int                   `json:"pool_id"`
	Subaccount       string                `json:"subaccount"`
	Owner            string                `json:"owner"`
	BalanceWeightX18 string                `json:"balance_weight_x18"`
	SubaccountInfo   NLPPoolSubaccountInfo `json:"subaccount_info"`
	OpenOrders       []OpenOrder           `json:"open_orders"`
}

type OpenOrder struct {
	ProductId      int    `json:"product_id"`
	Sender         string `json:"sender"`
	PriceX18       string `json:"price_x18"`
	Amount         string `json:"amount"`
	Expiration     string `json:"expiration"`
	OrderType      string `json:"order_type"`
	Nonce          string `json:"nonce"`
	Appendix       string `json:"appendix"`
	UnfilledAmount string `json:"unfilled_amount"`
	Digest         string `json:"digest"`
	PlacedAt       int    `json:"placed_at"`
	Id             int    `json:"id"`
}

// NLPPoolInfoData wraps the list of pools
type NLPPoolInfoData struct {
	NLPPools []NLPPool `json:"nlp_pools"`
}

// NLPPoolInfoResponse represents the response from GET /query?type=nlp_pool_info
type NLPPoolInfoResponse = BaseResponse[NLPPoolInfoData]

// ==========================================
// NLP Locked Balances Types
// ==========================================

// NLPBalanceDetail contains specific fields for NLP balances
type NLPBalanceDetail struct {
	Amount                   string `json:"amount"`
	LastCumulativeFundingX18 string `json:"last_cumulative_funding_x18"`
}

// NLPBalanceItem wraps the detail with a product ID
type NLPBalanceItem struct {
	ProductID int              `json:"product_id"`
	Balance   NLPBalanceDetail `json:"balance"`
}

// LockedBalanceEntry represents a specific lock entry with an unlock timestamp
type LockedBalanceEntry struct {
	Balance    NLPBalanceItem `json:"balance"`
	UnlockedAt string         `json:"unlocked_at"` // Timestamp string
}

// NLPLockedBalancesData represents the full response data
type NLPLockedBalancesData struct {
	BalanceLocked   NLPBalanceItem       `json:"balance_locked"`
	BalanceUnlocked NLPBalanceItem       `json:"balance_unlocked"`
	LockedBalances  []LockedBalanceEntry `json:"locked_balances"`
}

// NLPLockedBalancesResponse represents the response from GET /query?type=nlp_locked_balances
type NLPLockedBalancesResponse = BaseResponse[NLPLockedBalancesData]

// FeeRatesData contains fee schedules and sequencer fees.
// The array fields (like TakerFeeRatesX18) are typically indexed by Product ID.
type FeeRatesData struct {
	TakerFeeRatesX18        []string `json:"taker_fee_rates_x18"`
	MakerFeeRatesX18        []string `json:"maker_fee_rates_x18"`
	LiquidationSequencerFee string   `json:"liquidation_sequencer_fee"`
	HealthCheckSequencerFee string   `json:"health_check_sequencer_fee"`
	TakerSequencerFee       string   `json:"taker_sequencer_fee"`
	WithdrawSequencerFees   []string `json:"withdraw_sequencer_fees"`
}

// FeeRatesResponse represents the response from GET /query?type=fee_rates
type FeeRatesResponse = BaseResponse[FeeRatesData]

// HealthGroupsData contains the list of health groups.
// Each group is a list of product IDs (int) that are correlated for health calculation.
type HealthGroupsData struct {
	HealthGroups [][]int `json:"health_groups"`
}

// HealthGroupsResponse represents the response from GET /query?type=health_groups
type HealthGroupsResponse = BaseResponse[HealthGroupsData]

// LinkedSignerData contains the address of the linked signer.
// If no signer is linked, it often returns the zero address.
type LinkedSignerData struct {
	LinkedSigner string `json:"linked_signer"`
}

// LinkedSignerResponse represents the response from GET /query?type=linked_signer
type LinkedSignerResponse = BaseResponse[LinkedSignerData]

// InsuranceData contains the balance of the insurance fund
type InsuranceData struct {
	Insurance string `json:"insurance"` // int128 string
}

// InsuranceResponse represents the response from GET /query?type=insurance
type InsuranceResponse = BaseResponse[InsuranceData]
