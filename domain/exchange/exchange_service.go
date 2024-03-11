package exchange

import "time"

type ExchangeService interface {
	GetRate(date time.Time, currency string) (float64, error)
}
