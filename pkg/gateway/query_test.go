package gateway

import (
	"context"
	"fmt"
	"nado/pkg/types"
	"testing"
	"time"
)

const (
	TestSubaccount = "0xeae27ae6412147ed6d5692fd91709dad6dbfc342000000000000000000000000"

	TestSender = "0xeae27ae6412147ed6d5692fd91709dad6dbfc342000000000000000000000000"

	TestProductID = 1
)

func TestAllQueryMethods(t *testing.T) {
	client := setupTestClient(t)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	t.Run("GetSystemStatus", func(t *testing.T) {
		resp, err := client.GetSystemStatus(ctx)
		if err != nil {
			t.Fatalf("API Error: %v", err)
		}
		t.Logf("Success: Status=%s", resp.Data)
	})

	t.Run("GetContracts", func(t *testing.T) {
		resp, err := client.GetContracts(ctx)
		if err != nil {
			t.Fatalf("API Error: %v", err)
		}
		t.Logf("Success: Contracts=%s ChainID=%s", resp.Data.EndpointAddr, resp.Data.ChainID)
	})

	t.Run("GetFeeRates", func(t *testing.T) {
		resp, err := client.GetFeeRates(ctx, TestSender)
		if err != nil {
			t.Fatalf("API Error: %v", err)
		}
		t.Logf("Success: TakerFeeRates len=%d", len(resp.Data.TakerFeeRatesX18))
	})

	t.Run("GetSubaccountInfo", func(t *testing.T) {
		simulations := []types.SimulateTx{
			{
				ApplyDelta: &types.ApplyDeltaParams{
					ProductID:   2,
					Subaccount:  "0xeae27ae6412147ed6d5692fd91709dad6dbfc34264656661756c740000000000",
					AmountDelta: "100000000000000",
					VQuoteDelta: "3033500000000000000000",
				},
			},
		}

		ctx := context.Background()
		resp, err := client.GetSubaccountInfo(ctx, "0xeae27ae6412147ed6d5692fd91709dad6dbfc34264656661756c740000000000", simulations, true)

		if err != nil {
			t.Fatalf("API call failed: %v", err)
		}

		if !resp.IsSuccess() {
			t.Errorf("Expected success, got %s", resp.Status)
		}

		if resp.Data.PreState != nil {
			fmt.Println("fetch preState success")
		}

	})
	t.Run("GetHealthGroups", func(t *testing.T) {
		resp, err := client.GetHealthGroups(ctx)
		if err != nil {
			t.Fatalf("API Error: %v", err)
		}
		t.Logf("Success: Found %d health groups", len(resp.Data.HealthGroups))
	})

	t.Run("GetInsurance", func(t *testing.T) {
		resp, err := client.GetInsurance(ctx)
		if err != nil {
			t.Fatalf("API Error: %v", err)
		}
		t.Logf("Success: Insurance Fund=%s", resp.Data.Insurance)
	})

	t.Run("GetAllProducts", func(t *testing.T) {
		resp, err := client.GetAllProducts(ctx)
		if err != nil {
			t.Fatalf("API Error: %v", err)
		}
		t.Logf("Success: Spot=%d, Perp=%d", len(resp.Data.SpotProducts), len(resp.Data.PerpProducts))
	})

	t.Run("GetEdgeAllProducts", func(t *testing.T) {
		resp, err := client.GetEdgeAllProducts(ctx)
		if err != nil {
			t.Fatalf("API Error: %v", err)
		}
		t.Logf("Success: Found network for %d chains", len(resp.Data.EdgeAllProducts))
	})

	t.Run("GetSymbols_GET", func(t *testing.T) {
		resp, err := client.GetSymbols(ctx, "spot", nil)
		if err != nil {
			t.Fatalf("API Error: %v", err)
		}
		t.Logf("Success (GET): Found %d symbols", len(resp.Data.Symbols))
	})

	t.Run("GetSymbols_POST", func(t *testing.T) {
		resp, err := client.GetSymbols(ctx, "", []int{TestProductID})
		if err != nil {
			t.Fatalf("API Error: %v", err)
		}
		t.Logf("Success (POST): Found %d symbols (Filtered)", len(resp.Data.Symbols))
	})

	t.Run("GetMarketLiquidity", func(t *testing.T) {
		resp, err := client.GetMarketLiquidity(ctx, TestProductID, 5)
		if err != nil {
			t.Fatalf("API Error: %v", err)
		}
		t.Logf("Success: Bids=%d, Asks=%d", len(resp.Data.Bids), len(resp.Data.Asks))
	})

	t.Run("GetMarketPrice", func(t *testing.T) {
		resp, err := client.GetMarketPrice(ctx, TestProductID)
		if err != nil {
			t.Fatalf("API Error: %v", err)
		}
		t.Logf("Success: Bid=%s, Ask=%s", resp.Data.BidX18, resp.Data.AskX18)
	})

	t.Run("GetIsolatedPositions", func(t *testing.T) {
		resp, err := client.GetIsolatedPositions(ctx, TestSubaccount)
		if err != nil {
			t.Fatalf("API Error: %v", err)
		}
		t.Logf("Success: Found %d isolated positions", len(resp.Data.IsolatedPositions))
	})

	t.Run("GetNonces", func(t *testing.T) {
		resp, err := client.GetNonces(ctx, "0xeae27ae6412147ed6d5692fd91709dad6dbfc342")
		if err != nil {
			t.Fatalf("API Error: %v", err)
		}
		t.Logf("Success: Next nonce=%s", resp.Data.TxNonce)
	})

	t.Run("GetLinkedSigner", func(t *testing.T) {
		resp, err := client.GetLinkedSigner(ctx, TestSubaccount)
		if err != nil {
			t.Fatalf("API Error: %v", err)
		}
		t.Logf("Success: Linked Signer=%s", resp.Data.LinkedSigner)
	})

	t.Run("GetOrder", func(t *testing.T) {

		resp, err := client.GetOrder(ctx, 2, "0xa6f8df8f4f8bd52b1c6246d4266dc676ae9db64cf4f0af32e3a6bdeaf42fed38")
		if err != nil {
			t.Logf("API returned error (expected if digest invalid): %v", err)
		} else {
			t.Logf("Success: Order Status=%s", resp.Data.OrderType)
		}
	})

	t.Run("GetMaxOrderSize", func(t *testing.T) {
		params := types.MaxOrderSizeParams{
			ProductID:    TestProductID,
			Sender:       TestSender,
			PriceX18:     "20000000000000000000000", // 20000 (Example Price)
			Direction:    "long",
			SpotLeverage: true,
			ReduceOnly:   false,
			Isolated:     false,
		}
		resp, err := client.GetMaxOrderSize(ctx, params)
		if err != nil {
			t.Fatalf("API Error: %v", err)
		}
		t.Logf("Success: Max Order Size=%s", resp.Data.MaxOrderSize)
	})

	t.Run("GetMaxWithdrawable", func(t *testing.T) {
		resp, err := client.GetMaxWithdrawable(ctx, TestProductID, TestSender, false)
		if err != nil {
			t.Fatalf("API Error: %v", err)
		}
		t.Logf("Success: Max Withdrawable=%s", resp.Data.MaxWithdrawable)
	})

	t.Run("GetMaxNLPMintable", func(t *testing.T) {
		resp, err := client.GetMaxNLPMintable(ctx, TestSender, false)
		if err != nil {
			t.Fatalf("API Error: %v", err)
		}
		t.Logf("Success: Max Mintable Quote=%s", resp.Data.MaxQuoteAmount)
	})

	t.Run("GetMaxNLPBurnable", func(t *testing.T) {
		resp, err := client.GetMaxNLPBurnable(ctx, TestSender)
		if err != nil {
			t.Fatalf("API Error: %v", err)
		}
		t.Logf("Success: Max Burnable NLP=%s", resp.Data.MaxNLPAmount)
	})

	t.Run("GetNLPPoolInfo", func(t *testing.T) {
		resp, err := client.GetNLPPoolInfo(ctx)
		if err != nil {
			t.Fatalf("API Error: %v", err)
		}
		t.Logf("Success: Found %d pools", len(resp.Data.NLPPools))
	})

	t.Run("GetNLPPockedBalances", func(t *testing.T) {
		resp, err := client.GetNLPLockedBalances(ctx, TestSubaccount)
		if err != nil {
			t.Fatalf("API Error: %v", err)
		}
		t.Logf("Success: Locked Balances Count=%d", len(resp.Data.LockedBalances))
	})
}
