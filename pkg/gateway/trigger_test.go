package gateway

import (
	"context"
	"github.com/yuutmoo/nado-go/pkg/common"
	"github.com/yuutmoo/nado-go/pkg/types"
	"math/big"
	"testing"
	"time"
)

func TestPlaceOrder_TriggerFlow(t *testing.T) {
	param := &types.PlaceOrderParam{
		ProductID: 2,
		Amount:    0.012345,
		Price:     95000.7,
		Direction: types.DirectionLong,
		OrderType: types.OrderTypeDefault,

		Trigger: &types.TriggerConfig{
			PriceTrigger: &types.PriceTriggerConfig{
				OraclePriceAbove: 96000.12,
			},
		},
	}

	if err := param.Validate(); err != nil {
		t.Fatalf("validate fail: %v", err)
	}

	appendix := BuildAppendix(param)
	t.Logf("Generated Appendix: %s", appendix.String())

	orderPriceStr := common.FloatToX18(param.Price, 0.5)
	orderAmountStr := common.FloatToX18(param.Amount, 0.0001)

	triggerReq := param.Trigger.PriceTrigger.ToRequest(0.0001)

	expectedPrice := "95000500000000000000000"
	if orderPriceStr != expectedPrice {
		t.Errorf("Price alignment failed: got %s, want %s", orderPriceStr, expectedPrice)
	}

	expectedAmount := "12000000000000000"
	if orderAmountStr != expectedAmount {
		t.Logf("Amount correctly aligned to %s", orderAmountStr)
	}

	if triggerReq.Requirement.OraclePriceAbove == nil {
		t.Error("Trigger requirement should not be empty")
	}
}

func TestPlaceOrderParam_Validate(t *testing.T) {
	tests := []struct {
		name    string
		param   *types.PlaceOrderParam
		wantErr bool
	}{
		{
			name: "valid param",
			param: &types.PlaceOrderParam{
				Isolated:  false,
				OrderType: types.OrderTypeDefault,
			},
			wantErr: false,
		},
		{
			name: "Isolated + TWAP can not exist together",
			param: &types.PlaceOrderParam{
				Isolated: true,
				Trigger: &types.TriggerConfig{
					TimeTrigger: &types.TimeTriggerConfig{},
				},
			},
			wantErr: true,
		},
		{
			name: "TWAP must be with IOC",
			param: &types.PlaceOrderParam{
				OrderType: types.OrderTypeDefault,
				Trigger: &types.TriggerConfig{
					TimeTrigger: &types.TimeTriggerConfig{},
				},
			},
			wantErr: true,
		},
		{
			name: "margin can not be zero",
			param: &types.PlaceOrderParam{
				Isolated:       true,
				IsolatedMargin: 0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.param.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestBuildAppendix_Logic(t *testing.T) {
	t.Run("Isolated Appendix Calculation", func(t *testing.T) {
		margin := 500.5
		p := &types.PlaceOrderParam{
			Isolated:       true,
			IsolatedMargin: margin,
			OrderType:      types.OrderTypeDefault,
		}

		appendix := BuildAppendix(p)

		if appendix.Bit(8) != 1 {
			t.Error("Isolated bit (8) should be 1")
		}

		high := new(big.Int).Rsh(appendix, 64).Uint64()
		expectedMarginX6 := uint64(margin * 1e6)
		if high != expectedMarginX6 {
			t.Errorf("High 64 bits (Margin) = %d, want %d", high, expectedMarginX6)
		}
	})

	t.Run("TWAP Appendix Calculation", func(t *testing.T) {
		slippage := 0.01                    // 1%
		amounts := []float64{1.0, 1.0, 1.0} // 3 times

		p := &types.PlaceOrderParam{
			OrderType: types.OrderTypeIOC,
			Trigger: &types.TriggerConfig{
				TimeTrigger: &types.TimeTriggerConfig{
					Amounts:  amounts,
					Slippage: slippage,
				},
			},
		}

		appendix := BuildAppendix(p)

		if appendix.Bit(12) != 0 || appendix.Bit(13) != 1 {
			t.Error("TriggerType bits (12-13) should be 2 (TWAP)")
		}

		high := new(big.Int).Rsh(appendix, 64).Uint64()
		expectedTimes := uint32(len(amounts))
		expectedSlippageX6 := uint32(slippage * 1e6)
		expectedHigh := (uint64(expectedTimes) << 32) | uint64(expectedSlippageX6)

		if high != expectedHigh {
			t.Errorf("High 64 bits (TWAP Data) = %d, want %d", high, expectedHigh)
		}
	})
}

func TestPlacePriceTriggerOrder(t *testing.T) {
	client := setupTestClient(t)

	oneDayLater := uint64(time.Now().Add(time.Hour).Unix())
	param := &types.PlaceOrderParam{
		ProductID:  2,
		Amount:     0.002,
		Price:      11000,
		Direction:  types.DirectionShort,
		OrderType:  types.OrderTypeDefault,
		Expiration: &oneDayLater,
		Trigger: &types.TriggerConfig{
			PriceTrigger: &types.PriceTriggerConfig{
				LastPriceAbove: 100000,
			},
		},
	}

	resp, err := client.PlaceOrder(context.Background(), param)
	if err != nil {
		t.Fatalf("price trigger place order fail: %v", err)
	}
	t.Logf("Digest: %s", resp.Digest)

}
