package model

import "time"

var (
	OrderStatusFilled   = "FILLED"
	OrderStatusCanceled = "CANCELED"
)


type TradeStatus int

func (ts TradeStatus) String() string {
	return tradeStatusSymbol[ts]
}

var tradeStatusSymbol = [...]string{"UNFINISH", "PART_FINISH", "FINISH", "CANCEL", "REJECT", "CANCEL_ING", "FAIL"}

const (
	ORDER_UNFINISH TradeStatus = iota
	ORDER_PART_FINISH
	ORDER_FINISH
	ORDER_CANCEL
	ORDER_REJECT
	ORDER_CANCEL_ING
	ORDER_FAIL
)

const (
	OPEN_BUY   = "OB" //开多
	OPEN_SELL  = "OS" //开空
	CLOSE_BUY  = "CB" //平多
	CLOSE_SELL = "CS" //平空
)

type Order struct {
	ClientOrderID string
	OrderID       string
	Symbol        CurrencyPair
	Price         float64
	Amount        float64
	AvgPrice      float64
	DealAmount    float64
	Fee           float64
	Status        TradeStatus
	Side          TradeSide
	Offset        string
	OrderType     int   //0:default,1:maker,2:fok,3:ioc
	OrderTime     int64 // create  timestamp
	FinishedTime  int64 //finished timestamp
	//策略委托单
	LeverRate    float64 //倍数
	TriggerPrice float64
	AlgoType     int //1:限价 2:市场价；触发价格类型，默认是限价；为市场价时，委托价格不必填；

	Timestamp int64
	Date      time.Time
}
