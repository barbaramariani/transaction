package usecases

import (
	"log"

	"transaction/domain/entities"
	"transaction/domain/repositories"
)

type StoreTransaction struct {
	TransactionRepo repositories.TransactionRepository
}

func (s *StoreTransaction) Execute(t entities.Transaction) (string, error) {
	err := s.TransactionRepo.Store(t)
	if err != nil {
		return "", err
	}
	log.Printf("Transaction stored successfully - ID: %s, Description: %s, Amount: %f", t.ID, t.Description, t.Amount)
	return t.ID, nil
}
