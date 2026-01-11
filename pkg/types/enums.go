package types

type OrderType int

const (
	OrderTypeDefault  OrderType = 0
	OrderTypeIOC      OrderType = 1
	OrderTypeFOK      OrderType = 2
	OrderTypePostOnly OrderType = 3
)

type TriggerType int

const (
	TriggerTypeNone         TriggerType = 0
	TriggerTypePrice        TriggerType = 1
	TriggerTypeTwap         TriggerType = 2
	TriggerTypeCustomAmount TriggerType = 3
)

type OrderDirection string

const (
	DirectionLong  OrderDirection = "long"
	DirectionShort OrderDirection = "short"
)
