package infrastructure_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"transaction/infrastructure"

	"github.com/stretchr/testify/assert"
)

func TestFiscalDataTreasuryAPI_GetRate_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"data": [{"exchange_rate": "1.2"}]}`))
	}))
	defer server.Close()

	fiscalDataTreasuryAPI := infrastructure.NewFiscalDataTreasuryAPI(server.URL)
	date := time.Now()
	currency := "USD"

	rate, err := fiscalDataTreasuryAPI.GetRate(date, currency)

	assert.NoError(t, err)
	assert.Equal(t, 1.2, rate)
}

func TestFiscalDataTreasuryAPI_GetRate_Error(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`{"data": []}`))
	}))
	defer server.Close()

	fiscalDataTreasuryAPI := infrastructure.NewFiscalDataTreasuryAPI(server.URL)
	date := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	currency := "InvalidCurrency"

	rate, err := fiscalDataTreasuryAPI.GetRate(date, currency)

	assert.Error(t, err)
	assert.Zero(t, rate)
}
