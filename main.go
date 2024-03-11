package main

import (
	"net/http"
	"transaction/presentation"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/transactions/{id}", presentation.HandleRetrieveTransaction).Methods("GET")
	r.HandleFunc("/transactions", presentation.HandleStoreTransaction).Methods("POST")

	http.ListenAndServe(":8080", r)
}
