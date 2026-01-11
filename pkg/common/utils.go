package common

import (
	"github.com/shopspring/decimal"
	"github.com/yuutmoo/nado-go/pkg/types"
	"math/big"
	"time"
)

// GenerateNonce creates a timestamp-based nonce
// Logic: (CurrentTimeMs + 10s) << 20 + 1000
func GenerateNonce() uint64 {
	nowMs := time.Now().UnixMilli()
	recvTime := uint64(nowMs + 10000)
	return (recvTime << 20) + 1000
}

// FloatToX18 converts a float to a 1e18 string with tick size rounding
// Requires 'decimal' library
func FloatToX18(val float64, tickSize float64) string {
	dVal := decimal.NewFromFloat(val)

	if tickSize > 0 {
		dTick := decimal.NewFromFloat(tickSize)
		// Round down to nearest tick multiple: floor(val / tick) * tick
		dVal = dVal.Div(dTick).Floor().Mul(dTick)
	}

	dX18 := dVal.Shift(18)
	return dX18.Round(0).String()
}

// X18ToFloat converts a 1e18 string back to float64
func X18ToFloat(x18Str string) float64 {
	d, err := decimal.NewFromString(x18Str)
	if err != nil {
		return 0
	}
	f, _ := d.Shift(-18).Float64()
	return f
}

func BuildAppendix(params *types.PlaceOrderParam) *big.Int {
	// Bit Layout (from docs):
	// Version (0-7): 1
	// Isolated (8): 0 (assuming cross margin)
	// Order Type (9-10)
	// Reduce Only (11)

	opts := params.Options
	if opts == nil {
		opts = &types.OrderOptions{}
	}

	appendix := big.NewInt(0)

	// 1. Version (0-7): 1
	appendix.Or(appendix, big.NewInt(1))

	// 2. Isolated (8)
	if opts.Isolated != nil {
		appendix.SetBit(appendix, 8, 1)
	}

	// 3. Order Type (9-10)
	typeBig := big.NewInt(int64(params.OrderType))
	appendix.Or(appendix, typeBig.Lsh(typeBig, 9))

	// 4. Reduce Only (11)
	if params.ReduceOnly {
		appendix.SetBit(appendix, 11, 1)
	}

	// 5. Trigger Type (12-13)
	trigType := types.TriggerTypeNone
	if opts.Trigger != nil {
		trigType = opts.Trigger.Type
	}
	trigVal := big.NewInt(int64(trigType))
	appendix.Or(appendix, trigVal.Lsh(trigVal, 12))

	// 6. Value (64-127)
	valuePart := big.NewInt(0)

	if trigType == types.TriggerTypeTwap || trigType == types.TriggerTypeTwapCustomAmounts {
		if opts.Trigger != nil {
			slippageX6 := ToX6(opts.Trigger.TwapSlippage)
			valuePart.SetInt64(slippageX6)

			timesBig := big.NewInt(int64(opts.Trigger.TwapCount))
			valuePart.Or(valuePart, timesBig.Lsh(timesBig, 32))
		}
	} else if opts.Isolated != nil {
		marginX6 := ToX6(opts.Isolated.Amount)
		valuePart.SetInt64(marginX6)
	}

	appendix.Or(appendix, valuePart.Lsh(valuePart, 64))

	return appendix

}

// ToX6 converts a float to an integer with 1e6 precision (used for margins/slippage)
// Reference: Python `to_x6`
func ToX6(val float64) int64 {
	return decimal.NewFromFloat(val).Mul(decimal.NewFromFloat(1e6)).IntPart()
}

// FloatToIntString converts a float to a string integer with specific decimals
// e.g. val=100.5, decimals=6 (USDT) -> "100500000"
func FloatToIntString(val float64, decimals int32) string {
	d := decimal.NewFromFloat(val)
	mul := decimal.New(1, decimals)
	result := d.Mul(mul).Floor()
	return result.String()
}
