package entities

import "time"

type ConvertedTransaction struct {
	ID              string
	Description     string
	TransactionDate time.Time
	OriginalAmount  float64
	ExchangeRate    float64
	ConvertedAmount float64
}
