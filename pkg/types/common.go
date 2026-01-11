package types

import "fmt"

// BaseResponse is the standard wrapper for Nado API responses
type BaseResponse[T any] struct {
	Status      string `json:"status"`
	Data        T      `json:"data"`
	RequestType string `json:"request_type"`

	ErrorCode int    `json:"error_code,omitempty"`
	ErrorMsg  string `json:"error,omitempty"`
}

func (r *BaseResponse[T]) IsSuccess() bool {
	return r.Status == "success"
}

func (r *BaseResponse[T]) Error() string {
	if r.IsSuccess() {
		return ""
	}
	return fmt.Sprintf("api error (%s): [%d] %s", r.Status, r.ErrorCode, r.ErrorMsg)
}

type ExecuteResponse[T any] struct {
	BaseResponse[T]

	Signature string `json:"signature"`
	ID        int64  `json:"id"`
}

type ProductInfo struct {
	ProductID  int     `json:"product_id"`
	PriceTick  float64 `json:"price_tick"`
	AmountTick float64 `json:"amount_tick"`
	MinSize    float64 `json:"min_size"`
}
