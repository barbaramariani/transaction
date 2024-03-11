package presentation_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"transaction/presentation"
)

func TestHandleStoreTransaction_WhenPostingATransaction_CheckStatusCreatedAndResponse_Success(t *testing.T) {
	req, err := http.NewRequest("POST", "/transactions", bytes.NewBuffer([]byte(`{"description": "test", "transaction_date": "2022-01-01", "purchase_amount_usd": 100}`)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(presentation.HandleStoreTransaction)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	if !strings.Contains(rr.Body.String(), `"description":"test"`) || !strings.Contains(rr.Body.String(), `"transaction_date":"2022-01-01T00:00:00Z"`) || !strings.Contains(rr.Body.String(), `"purchase_amount_usd":100`) {
		t.Errorf("handler returned unexpected body: got %v", rr.Body.String())
	}
}

func TestHandleStoreTransaction_WhenPostingATransactionWithInvalidInput_CheckStatusBadRequestAndResponse_ValidationErrors(t *testing.T) {
	req, err := http.NewRequest("POST", "/transactions", bytes.NewBuffer([]byte(`{"description": "this description is too long and should not be accepted", "transaction_date": "2022-01-01", "purchase_amount_usd": 100}`)))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(presentation.HandleStoreTransaction)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	expected := `Validation errors: Description: must be less than or equal to 50 characters,`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHandleRetrieveTransaction_WhenTransactionDoesNotExist_CheckStatusNotFoundAndResponse_Error(t *testing.T) {

	req, err := http.NewRequest("GET", "/transactions/non-existent-id", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(presentation.HandleRetrieveTransaction)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}

	expected := `Error retrieving transaction: transaction not found`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
