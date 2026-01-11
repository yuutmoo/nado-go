package common

import (
	"github.com/shopspring/decimal"
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
