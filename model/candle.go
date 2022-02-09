package model

import (
	"fmt"
	"time"

	"github.com/go-sohunjug/logger"
	"github.com/mitchellh/mapstructure"
)

var Log = logger.NewHelper(logger.With(DefaultLogger, "module", "quant/model"))

type Candle struct {
	Symbol    CurrencyPair
	Timestamp int64
	Interval  string
	Open      float64
	High      float64
	Low       float64
	Close     float64
	BaseVol   float64
	QuoteVol  float64
	Date      time.Time
}

func (c Candle) Datetime() time.Time {
	return c.Date
}

func (c Candle) String() string {
	return fmt.Sprintf("timestamp %d, %s open:%f close:%f low:%f high:%f quotvol:%f basevol:%f", c.Timestamp, c.Datetime().String(), c.Open, c.Close, c.Low, c.High, c.QuoteVol, c.BaseVol)
}

// CandleList candle list
type CandleList []*Candle

// Merge merge multi candle to one
func (l CandleList) Merge() (ret *Candle) {
	if len(l) == 0 {
		return
	}
	ret = new(Candle)
	ret.Timestamp = l[0].Timestamp
	ret.Open = l[0].Open
	ret.High = l.High()
	ret.Low = l.Low()
	ret.Close = l[len(l)-1].Close
	for _, v := range l {
		ret.QuoteVol += v.QuoteVol
		ret.BaseVol += v.BaseVol
	}
	return
}

func (l CandleList) High() (ret float64) {
	for _, v := range l {
		if ret < v.High {
			ret = v.High
		}
	}
	return
}

func (l CandleList) Low() (ret float64) {
	for _, v := range l {
		if ret == 0 {
			ret = v.Low
			continue
		}
		if ret > v.Low {
			ret = v.Low
		}
	}
	return
}

// Map2Candle convert candle to map
func Map2Candle(data interface{}) (candle *Candle) {
	candle, ok := data.(*Candle)
	if ok {
		return
	}
	candle = new(Candle)
	err := mapstructure.Decode(data, &candle)
	if err != nil {
		Log.Error("Map2Candle failed:", data, err.Error())
	}
	return
}

type CandleFn func(candle Candle)

type Param struct {
	Name string
	Type string
	Info string
}

type ParamData map[string]interface{}

func (d ParamData) GetString(key, defaultValue string) string {
	v, ok := d[key]
	if !ok {
		return defaultValue
	}
	ret := v.(string)
	if ret == "" {
		return defaultValue
	}
	return ret
}
func (d ParamData) GetFloat(key string, defaultValue float64) float64 {
	v, ok := d[key]
	if !ok {
		return defaultValue
	}
	ret := v.(float64)
	if ret == 0 {
		return defaultValue
	}
	return ret
}
