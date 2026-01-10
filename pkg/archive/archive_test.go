package archive

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func getTestClient() *ArchiveClient {
	a := NewArchiveClient()
	return a
}

func TestIntegration_GetCandlesticks(t *testing.T) {
	client := getTestClient()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	limit := 5
	params := CandlestickParams{
		ProductId:   1,
		Granularity: 3600,
		Limit:       &limit,
	}

	candles, err := client.GetCandlesticks(ctx, params)

	assert.NoError(t, err)
	assert.NotEmpty(t, candles)
	if len(candles) > 0 {
		assert.Equal(t, 1, candles[0].ProductId)
		fmt.Printf("Latest Candle Close: %s\n", candles[0].CloseX18)
	}
}

func TestIntegration_GetFundingRate(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	rate, err := client.GetFundingRate(ctx, 1)
	assert.NoError(t, err)
	assert.NotNil(t, rate)
	assert.NotEmpty(t, rate.FundingRateX18)

	rates, err := client.GetFundingRates(ctx, []int{1, 2})
	assert.NoError(t, err)
	assert.NotNil(t, rates)

	for id, r := range rates {
		fmt.Printf("Product %s Rate: %s\n", id, r.FundingRateX18)
	}
}

func TestIntegration_GetProductSnapshots(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	params := ProductSnapshotParams{
		ProductId: 1,
		Limit:     2,
	}

	resp, err := client.GetProductSnapshots(ctx, params)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.True(t, len(resp.Products) > 0)
	assert.True(t, len(resp.Txs) > 0)

	if len(resp.Products) > 0 && len(resp.Txs) > 0 {
		assert.Equal(t, resp.Products[0].SubmissionIdx, resp.Txs[0].SubmissionIdx)
		fmt.Printf("Snapshot Time: %s\n", resp.Txs[0].Timestamp)
	}
}

func TestIntegration_GetEvents(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	params := EventsParams{
		ProductIds: []int{1},
		EventTypes: []string{"match_orders"},
		Limit:      map[string]int{"raw": 5},
	}

	resp, err := client.GetEvents(ctx, params)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	if len(resp.Events) > 0 {
		assert.Equal(t, "match_orders", resp.Events[0].EventType)
	}
}

func TestIntegration_GetMatches(t *testing.T) {
	client := getTestClient()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	limit := 5
	params := MatchesParams{
		Limit: &limit,
	}

	resp, err := client.GetMatches(ctx, params)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	if len(resp.Matches) > 0 {
		match := resp.Matches[0]
		assert.NotEmpty(t, match.Digest)
		assert.NotEmpty(t, match.Order.PriceX18)

		foundTx := false
		for _, tx := range resp.Txs {
			if tx.SubmissionIdx == match.SubmissionIdx {
				assert.NotEmpty(t, tx.Timestamp)
				fmt.Printf("Match at %s, Price: %s\n", tx.Timestamp, match.Order.PriceX18)
				foundTx = true
				break
			}
		}
		assert.True(t, foundTx, "Every match should have a corresponding Tx entry")
	}
}

func TestIntegration_GetEdgeCandlesticks(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	limit := 5
	params := EdgeCandlestickParams{
		ProductId:   1, // BTC
		Granularity: 3600,
		Limit:       &limit,
	}

	candles, err := client.GetEdgeCandlesticks(ctx, params)

	assert.NoError(t, err)
	if len(candles) > 0 {
		assert.Equal(t, 1, candles[0].ProductId)
		assert.NotEmpty(t, candles[0].HighX18)
		fmt.Printf("Edge Candle - High: %s, Low: %s\n", candles[0].HighX18, candles[0].LowX18)
	}
}

func TestIntegration_GetEvents_LimitTypes(t *testing.T) {
	client := getTestClient()
	ctx := context.Background()

	params := EventsParams{
		ProductIds: []int{1},
		Limit:      map[string]int{"txs": 1},
	}

	resp, err := client.GetEvents(ctx, params)
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	if len(resp.Txs) > 0 {
		assert.LessOrEqual(t, len(resp.Txs), 1)
		fmt.Printf("Events received for 1 transaction: %d\n", len(resp.Events))
	}
}

func TestIntegration_GetMatchedOrders(t *testing.T) {
	client := getTestClient()
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	t.Run("QueryBySubaccount", func(t *testing.T) {
		params := MatchedOrdersParams{
			Subaccounts: []string{"0x12a0b4888021576eb10a67616dd3dd3d9ce206b664656661756c740000000000"},
			ProductIds:  []int{1, 2},
			Limit:       5,
		}

		orders, err := client.GetMatchedOrders(ctx, params)
		assert.NoError(t, err)

		if len(orders) > 0 {
			assert.NotEmpty(t, orders[0].Digest)
			assert.NotEmpty(t, orders[0].SubmissionIdx)
			fmt.Printf("Order Found: %s, Base Filled: %s\n", orders[0].Digest, orders[0].BaseFilled)
		}
	})

	t.Run("ValidationConflict", func(t *testing.T) {
		params := MatchedOrdersParams{
			Subaccounts: []string{"0x1"},
			Digests:     []string{"0xabc"},
		}
		_, err := client.GetMatchedOrders(ctx, params)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "cannot provide both 'subaccounts' and 'digests'")
	})

	t.Run("IsolatedFilter", func(t *testing.T) {
		isIsolated := true
		params := MatchedOrdersParams{
			Subaccounts: []string{"0x12a0b4888021576eb10a67616dd3dd3d9ce206b664656661756c740000000000"},
			Isolated:    &isIsolated,
			Limit:       1,
		}

		_, err := client.GetMatchedOrders(ctx, params)
		assert.NoError(t, err)
	})
}
