package entities

import (
	"encoding/json"
	"math"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID              string    `json:"id"`
	Description     string    `json:"description"`
	TransactionDate time.Time `json:"transaction_date"`
	Amount          float64   `json:"purchase_amount_usd"`
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	type transactionAlias Transaction
	var tmp struct {
		*transactionAlias
		ID              string `json:"id"`
		TransactionDate string `json:"transaction_date"`
	}

	tmp.transactionAlias = (*transactionAlias)(t)

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	if tmp.ID == "" {
		tmp.ID = uuid.New().String()
	}

	transactionDate, err := time.Parse("2006-01-02", tmp.TransactionDate)
	if err != nil {
		return err
	}
	t.TransactionDate = transactionDate

	t.ID = tmp.ID

	t.Amount = math.Round(t.Amount*100) / 100

	return nil
}
