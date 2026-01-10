package types

type OrderType int

const (
	OrderTypeDefault  OrderType = 0
	OrderTypeIOC      OrderType = 1
	OrderTypeFOK      OrderType = 2
	OrderTypePostOnly OrderType = 3
)

// 新增：触发类型 (Trigger Type)
type TriggerType int

const (
	TriggerTypeNone              TriggerType = 0
	TriggerTypePrice             TriggerType = 1 // 止盈止损通常用这个
	TriggerTypeTwap              TriggerType = 2
	TriggerTypeTwapCustomAmounts TriggerType = 3
)

type OrderDirection string

const (
	DirectionLong  OrderDirection = "long"
	DirectionShort OrderDirection = "short"
)
