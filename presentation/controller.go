package presentation

import (
	"encoding/json"
	"io"
	"net/http"
	"transaction/application/usecases"
	"transaction/domain/entities"
	"transaction/domain/validators"

	"github.com/gorilla/mux"
)

func HandleStoreTransaction(w http.ResponseWriter, r *http.Request, uc usecases.StoreTransaction) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var t entities.Transaction
	err = json.Unmarshal(body, &t)
	if err != nil {
		http.Error(w, "Error unmarshalling request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	v := validators.NewValidator()
	v.Add("Description", validators.ValidateLength(50))
	v.Add("Amount", validators.ValidatePositiveAmount())

	data := map[string]interface{}{
		"ID":              t.ID,
		"Description":     t.Description,
		"TransactionDate": t.TransactionDate,
		"Amount":          t.Amount,
	}

	errors := v.Validate(data)
	if len(errors) > 0 {
		errorMessage := "Validation errors: "
		for _, err := range errors {
			errorMessage += err.Error() + ", "
		}
		http.Error(w, errorMessage, http.StatusBadRequest)
		return
	}

	_, err = uc.Execute(t)
	if err != nil {
		http.Error(w, "Error storing transaction", http.StatusInternalServerError)
		return
	}
	jsonResponse, err := json.Marshal(t)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func HandleRetrieveTransaction(w http.ResponseWriter, r *http.Request, uc usecases.RetrieveTransaction) {

	vars := mux.Vars(r)
	id := vars["id"]
	currency := r.URL.Query().Get("currency")

	convertedTransaction, err := uc.Execute(id, currency)
	if err != nil {
		errorMessage := "Error retrieving transaction: " + err.Error()
		http.Error(w, errorMessage, http.StatusNotFound)
		return
	}

	jsonResponse, err := json.Marshal(convertedTransaction)
	if err != nil {
		http.Error(w, "Error marshalling response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
