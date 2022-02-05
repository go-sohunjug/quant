package model

import (
	"fmt"
	"reflect"
	"time"

	"github.com/go-sohunjug/quant/indicator"
)

// Events
const (
	EventCandleParam    = "event.candle_param"
	EventCandle         = "event.candle"
	EventTicker         = "event.ticker"
	EventOrder          = "event.order"
	EventOrderCancelAll = "event.order_cancel_all"
	// own trades
	EventTrade       = "event.trade"
	EventPosition    = "event.position"
	EventCurPosition = "event.cur_position" // position of current script
	EventRiskLimit   = "event.risk_limit"
	EventDepth       = "event.depth"
	// all trades in the markets
	EventTradeHistory = "event.trade_history"

	EventBalance     = "event.balance"
	EventBalanceInit = "event.balance_init"

	EventWatch = "event.watch"

	EventNotify = "event.notify"
)

var (
	EventTypes = map[string]reflect.Type{
		EventCandleParam: reflect.TypeOf(CandleParam{}),
		EventCandle:      reflect.TypeOf(Candle{}),
		EventTicker:       reflect.TypeOf(Ticker{}),
		EventOrder:       reflect.TypeOf(TradeAction{}),
		// EventOrderCancelAll     = "order_cancel_all"
		EventTrade:    reflect.TypeOf(Trade{}),
		EventPosition: reflect.TypeOf(Position{}),
		// EventCurPosition        = "cur_position" // position of current script
		// EventRiskLimit          = "risk_limit"
		EventDepth:        reflect.TypeOf(Depth{}),
		EventTradeHistory: reflect.TypeOf(Trade{}),
		EventBalance:      reflect.TypeOf(Balance{}),
		EventBalanceInit:  reflect.TypeOf(BalanceInfo{}),
		EventWatch:        reflect.TypeOf(WatchParam{}),

		EventNotify: reflect.TypeOf(NotifyEvent{}),
	}
)

type Engine interface {
	OpenLong(price, amount float64)
	CloseLong(price, amount float64)
	OpenShort(price, amount float64)
	CloseShort(price, amount float64)
	StopLong(price, amount float64)
	StopShort(price, amount float64)
	CancelAllOrder()
	AddIndicator(name string, params ...int) (ind indicator.CommonIndicator)
	Position() (pos, price float64)
	Balance() float64
	Log(v ...interface{})
	Watch(watchType string)
	SendNotify(content, contentType string)
	Merge(src, dst string, fn CandleFn)
	SetBalance(balance float64)

	// call for goscript
	UpdatePosition(pos, price float64)
	OnCandle(candle Candle)
	UpdateBalance(balance float64)
}

// CandleParam get candle param
type CandleParam struct {
	Start    time.Time
	End      time.Time
	Exchange string
	BinSize  string
	Symbol   string
}

// NotifyEvent event to send notify
type NotifyEvent struct {
	Type    string // text,markdown
	Content string
}

// RiskLimit risk limit
type RiskLimit struct {
	Code         string  // symbol info, empty = global
	Lever        float64 // lever
	MaxLostRatio float64 // max lose ratio
}

// Key key of r
func (r RiskLimit) Key() string {
	return fmt.Sprintf("%s-%.2f", r.Code, r.Lever)
}

// WatchParam add watch event param
type WatchParam struct {
	Type  string
	Param map[string]interface{}
}

// BalanceInfo balance
type BalanceInfo struct {
	Balance float64
}
