package model

import (
	"strings"
	"time"
)

// A->B(A兑换为B)
type CurrencyPair struct {
	CurrencyA      Currency
	CurrencyB      Currency
	Exchange       string
	ContractType   string
	DeliveryDate   time.Time
	AmountTickSize int // 下单量精度
	PriceTickSize  int //交易对价格精度
}

func (pair CurrencyPair) String() string {
	return pair.ToSymbol("/")
}

func (pair CurrencyPair) Eq(c2 CurrencyPair) bool {
	return pair.String() == c2.String()
}

func NewCurrencyPair(currencyA Currency, currencyB Currency) CurrencyPair {
	return CurrencyPair{CurrencyA: currencyA, CurrencyB: currencyB}
}

func NewCurrencyPairWithString(A string, B string) CurrencyPair {
	currencyA := NewCurrency(A, "")
	currencyB := NewCurrency(B, "")
	return CurrencyPair{CurrencyA: currencyA, CurrencyB: currencyB}
}

func NewCurrencyPairDefault(currencyPairSymbol string) CurrencyPair {
	pair := NewCurrencyPairSep(currencyPairSymbol, "/")
	if pair == UNKNOWN_PAIR {
		pair = NewCurrencyPairSep(currencyPairSymbol, "_")
	}
	return pair
}

func NewCurrencyPairSep(currencyPairSymbol string, sep string) CurrencyPair {
	currencys := strings.Split(currencyPairSymbol, sep)
	if len(currencys) >= 2 {
		return CurrencyPair{CurrencyA: NewCurrency(currencys[0], ""),
			CurrencyB: NewCurrency(currencys[1], "")}
	}
	return UNKNOWN_PAIR
}

func (pair *CurrencyPair) SetAmountTickSize(tickSize int) CurrencyPair {
	pair.AmountTickSize = tickSize
	return *pair
}

func (pair *CurrencyPair) SetPriceTickSize(tickSize int) CurrencyPair {
	pair.PriceTickSize = tickSize
	return *pair
}

func (pair CurrencyPair) ToSymbol(joinChar string) string {
	return strings.Join([]string{pair.CurrencyA.Symbol, pair.CurrencyB.Symbol}, joinChar)
}

func (pair CurrencyPair) ToSymbolReverse(joinChar string) string {
	return strings.Join([]string{pair.CurrencyB.Symbol, pair.CurrencyA.Symbol}, joinChar)
}

func (pair CurrencyPair) AdaptUsdtToUsd() CurrencyPair {
	CurrencyB := pair.CurrencyB
	if pair.CurrencyB.Eq(USDT) {
		CurrencyB = USD
	}
	pair.CurrencyB = CurrencyB
	return pair
}

func (pair CurrencyPair) AdaptUsdToUsdt() CurrencyPair {
	CurrencyB := pair.CurrencyB
	if pair.CurrencyB.Eq(USD) {
		CurrencyB = USDT
	}
	pair.CurrencyB = CurrencyB
	return pair
}

//for to symbol lower , Not practical '==' operation method
func (pair CurrencyPair) ToUpper() CurrencyPair {
	return CurrencyPair{CurrencyA: Currency{Symbol: strings.ToUpper(pair.CurrencyA.Symbol), Desc: pair.CurrencyA.Desc},
		CurrencyB: Currency{Symbol: strings.ToUpper(pair.CurrencyB.Symbol), Desc: pair.CurrencyB.Desc}}
}

func (pair CurrencyPair) ToLower() CurrencyPair {
	return CurrencyPair{CurrencyA: Currency{Symbol: strings.ToLower(pair.CurrencyA.Symbol), Desc: pair.CurrencyA.Desc},
		CurrencyB: Currency{Symbol: strings.ToLower(pair.CurrencyB.Symbol), Desc: pair.CurrencyB.Desc}}
}

func (pair CurrencyPair) Reverse() CurrencyPair {
	return CurrencyPair{CurrencyA: pair.CurrencyB, CurrencyB: pair.CurrencyA,
		AmountTickSize: pair.AmountTickSize, PriceTickSize: pair.PriceTickSize}
}
