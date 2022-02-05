package model

import "time"

type TradeType int

const (
	DirectLong  TradeType = 1
	DirectShort TradeType = 1 << 1

	Limit  TradeType = 1 << 3
	Market TradeType = 1 << 4
	Stop   TradeType = 1 << 5

	Open  TradeType = 1 << 6
	Close TradeType = 1 << 7

	OpenLong   = Open | DirectLong
	OpenShort  = Open | DirectShort
	CloseLong  = Close | DirectLong
	CloseShort = Close | DirectShort
	StopLong   = Stop | DirectLong
	StopShort  = Stop | DirectShort
)

type TradeSide int

const (
	BUY TradeSide = 1 + iota
	SELL
	BUY_MARKET
	SELL_MARKET
)

func (ts TradeSide) String() string {
	switch ts {
	case 1:
		return "BUY"
	case 2:
		return "SELL"
	case 3:
		return "BUY_MARKET"
	case 4:
		return "SELL_MARKET"
	default:
		return "UNKNOWN"
	}
}

type Trade struct {
	Symbol    CurrencyPair `json:"omitempty"`
	Tid       int64        `json:"tid"`
	Type      TradeSide    `json:"type"`
	Amount    float64      `json:"amount,string"`
	Price     float64      `json:"price,string"`
	Timestamp int64        `json:"date_ms"`
	Date      time.Time
	Side      string
	Remark    string
}

// TradeAction trade action
type TradeAction struct {
	Action    TradeType
	Amount    float64
	Price     float64
	Date      time.Time
	Timestamp int64
}

func (a TradeType) IsLong() bool {
	if a&OpenLong == OpenLong || a&CloseShort == CloseShort || a&StopShort == StopShort {
		return true
	}
	return false
}

func (a TradeType) IsOpen() bool {
	if a&Open == Open {
		return true
	}
	return false
}

func (a TradeType) IsStop() bool {
	if a&Stop == Stop {
		return true
	}
	return false
}

func (t TradeType) String() (ret string) {
	if t&Limit == Limit {
		ret += "Limit"
	} else if t&Market == Market {
		ret += "Market"
	} else if t&Stop == Stop {
		ret += "Stop"
	}
	if t&Open == Open {
		ret += "Open"
	} else if t&Close == Close {
		ret += "Close"
	}

	if t&DirectLong == DirectLong {
		ret += "Long"
	} else if t&DirectShort == DirectShort {
		ret += "Short"
	}
	return
}
