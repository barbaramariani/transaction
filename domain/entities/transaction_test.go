package entities_test

import (
	"testing"
	"time"

	"transaction/domain/entities"

	"github.com/stretchr/testify/assert"
)

func TestTransaction_UnmarshalJSON_WithouID_GeneratesId_Success(t *testing.T) {
	transaction := &entities.Transaction{}
	data := []byte(`{
		"description": "test",
		"transaction_date": "2022-01-01",
		"purchase_amount_usd": 100.0
	}`)

	err := transaction.UnmarshalJSON(data)

	assert.NoError(t, err)
	assert.NotEmpty(t, transaction.ID)
	assert.Equal(t, "test", transaction.Description)
	assert.Equal(t, time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC), transaction.TransactionDate)
	assert.Equal(t, 100.0, transaction.Amount)
}

func TestTransaction_UnmarshalJSON_Error(t *testing.T) {
	transaction := &entities.Transaction{}
	data := []byte(`{
		"id": "90445774-14b2-44e1-8caf-3afd9bffcb83",
		"description": "test",
		"transaction_date": "invalid_date",
		"purchase_amount_usd": 100.0
	}`)

	err := transaction.UnmarshalJSON(data)

	assert.Error(t, err)
}
