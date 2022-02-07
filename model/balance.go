package model

type Account struct {
	Currency  Currency
	Available float64
	Frozen    float64
	Balance   float64
}

