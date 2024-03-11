package usecases

import (
	"log"
	"math"

	"transaction/domain/entities"
	"transaction/domain/exchange"
	"transaction/domain/repositories"
)

type RetrieveTransaction struct {
	TransactionRepo repositories.TransactionRepository
	ExchangeService exchange.ExchangeService
}

func (r *RetrieveTransaction) Execute(id string, currency string) (entities.ConvertedTransaction, error) {
	log.Printf("Retrieving transaction with ID: %s", id)
	transaction, err := r.TransactionRepo.Retrieve(id)
	if err != nil {
		log.Printf("Error retrieving transaction: %v", err)
		return entities.ConvertedTransaction{}, err
	}

	log.Printf("Getting exchange rate for transaction date: %v and currency: %s", transaction.TransactionDate, currency)
	rate, err := r.ExchangeService.GetRate(transaction.TransactionDate, currency)
	if err != nil {
		log.Printf("Error getting exchange rate: %v", err)
		return entities.ConvertedTransaction{}, err
	}

	convertedAmount := math.Round(transaction.Amount*rate*100) / 100

	return entities.ConvertedTransaction{
		ID:              transaction.ID,
		Description:     transaction.Description,
		TransactionDate: transaction.TransactionDate,
		OriginalAmount:  transaction.Amount,
		ExchangeRate:    rate,
		ConvertedAmount: convertedAmount,
	}, nil
}
