package infrastructure

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type FiscalDataTreasuryAPI struct {
	baseURL string
}

func NewFiscalDataTreasuryAPI(baseURL string) *FiscalDataTreasuryAPI {
	return &FiscalDataTreasuryAPI{baseURL: baseURL}
}

func (f *FiscalDataTreasuryAPI) GetRate(date time.Time, currency string) (float64, error) {
	sixMonthsAgo := time.Now().AddDate(0, -6, 0)
	formattedDate := sixMonthsAgo.Format("2006-01-02")
	url := fmt.Sprintf("%s/services/api/fiscal_service/v1/accounting/od/rates_of_exchange?fields=exchange_rate,record_date&filter=currency:eq:%s,record_date:gte:%s&sort=-record_date&format=json", f.baseURL, currency, formattedDate)
	log.Printf("Sending GET request to URL: %s", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error occurred while sending GET request: %s", err.Error())
	}
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	if data["data"] == nil || len(data["data"].([]interface{})) == 0 {
		return 0, fmt.Errorf("no exchange rate available for currency %s within the last 6 months", currency)
	}

	exchangeRate := data["data"].([]interface{})[0].(map[string]interface{})["exchange_rate"].(string)
	rate, err := strconv.ParseFloat(exchangeRate, 64)
	if err != nil {
		return 0, err
	}

	return rate, nil
}
