package main

import (
	"log"
	"net/http"
	"transaction/application/usecases"
	"transaction/presentation"

	"transaction/infrastructure"

	"github.com/gorilla/mux"
)

func main() {
	db, err := infrastructure.InitDB(":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	repo := infrastructure.NewSqliteTransactionRepository(db)
	r := mux.NewRouter()
	retrieveTransaction := usecases.RetrieveTransaction{
		TransactionRepo: repo,
		ExchangeService: infrastructure.NewFiscalDataTreasuryAPI("https://api.fiscaldata.treasury.gov"),
	}
	storeTransaction := usecases.StoreTransaction{TransactionRepo: repo}

	r.HandleFunc("/transactions/{id}", func(w http.ResponseWriter, r *http.Request) {
		presentation.HandleRetrieveTransaction(w, r, retrieveTransaction)
	}).Methods("GET")
	r.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		presentation.HandleStoreTransaction(w, r, storeTransaction)
	}).Methods("POST")

	http.ListenAndServe(":8080", r)
}
