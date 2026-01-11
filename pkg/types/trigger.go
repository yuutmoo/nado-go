package types

import (
	"fmt"
	"github.com/yuutmoo/nado-go/pkg/common"
)

// -----------------------trigger---------------------------------------

type PriceRequirement struct {
	OraclePriceAbove *string `json:"oracle_price_above,omitempty"`
	OraclePriceBelow *string `json:"oracle_price_below,omitempty"`
	LastPriceAbove   *string `json:"last_price_above,omitempty"`
	LastPriceBelow   *string `json:"last_price_below,omitempty"`
	MidPriceAbove    *string `json:"mid_price_above,omitempty"`
	MidPriceBelow    *string `json:"mid_price_below,omitempty"`
}

type TriggerDependency struct {
	Digest        string `json:"digest"`
	OnPartialFill bool   `json:"on_partial_fill"`
}

type TriggerConfig struct {
	PriceTrigger *PriceTriggerConfig
	TimeTrigger  *TimeTriggerConfig
}

type TriggerCriteria struct {
	PriceTrigger *PriceTriggerRequest `json:"price_trigger,omitempty"`
	TimeTrigger  *TimeTriggerRequest  `json:"time_trigger,omitempty"`
}

type PriceTriggerRequest struct {
	Requirement PriceRequirement   `json:"price_requirement"`
	Dependency  *TriggerDependency `json:"dependency,omitempty"`
}
type TimeTriggerRequest struct {
	Interval int64    `json:"interval"`
	Amounts  []string `json:"amounts,omitempty"`
}

type TimeTriggerConfig struct {
	Interval int64     `json:"interval"`
	Amounts  []float64 `json:"amounts,omitempty"`
	Times    int       `json:"-"`
	Slippage float64   `json:"-"`
}

type PriceTriggerConfig struct {
	OraclePriceAbove float64
	OraclePriceBelow float64
	LastPriceAbove   float64
	LastPriceBelow   float64
	MidPriceAbove    float64
	MidPriceBelow    float64

	Dependency *TriggerDependency
}

func (c *PriceTriggerConfig) ToRequest(tick float64) *PriceTriggerRequest {
	ptr := func(val float64) *string {
		if val <= 0 {
			return nil
		}
		s := common.FloatToX18(val, tick)
		return &s
	}
	return &PriceTriggerRequest{
		Requirement: PriceRequirement{
			OraclePriceAbove: ptr(c.OraclePriceAbove),
			OraclePriceBelow: ptr(c.OraclePriceBelow),
			LastPriceAbove:   ptr(c.LastPriceAbove),
			LastPriceBelow:   ptr(c.LastPriceBelow),
			MidPriceAbove:    ptr(c.MidPriceAbove),
			MidPriceBelow:    ptr(c.MidPriceBelow),
		},
		Dependency: c.Dependency,
	}
}

func (c *TimeTriggerConfig) ToRequest(tick float64) *TimeTriggerRequest {
	var amountsStr []string
	for _, a := range c.Amounts {
		amountsStr = append(amountsStr, common.FloatToX18(a, tick))
	}
	return &TimeTriggerRequest{
		Interval: c.Interval,
		Amounts:  amountsStr,
	}
}

func (p *PlaceOrderParam) Validate() error {

	if p.Isolated && p.Trigger != nil && p.Trigger.TimeTrigger != nil {
		return fmt.Errorf("isolated margin cannot be combined with TWAP orders")
	}

	if p.Trigger != nil && p.Trigger.TimeTrigger != nil {
		if p.OrderType != OrderTypeIOC {
			return fmt.Errorf("TWAP orders must use IOC (Immediate or Cancel) execution type")
		}
	}

	if p.Isolated && p.IsolatedMargin <= 0 {
		return fmt.Errorf("isolated margin must be greater than 0 when isolated mode is active")
	}

	if p.Trigger != nil && p.Trigger.PriceTrigger != nil {
		req := p.Trigger.PriceTrigger
		if req.OraclePriceAbove == 0 && req.OraclePriceBelow == 0 &&
			req.LastPriceAbove == 0 && req.LastPriceBelow == 0 &&
			req.MidPriceAbove == 0 && req.MidPriceBelow == 0 {
			return fmt.Errorf("price trigger enabled but no price requirement (above/below) provided")
		}
	}

	return nil
}
