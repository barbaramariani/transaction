package infrastructure

import (
	"errors"
	"log"
	"sync"
	"transaction/domain/entities"
)

type InMemoryTransactionRepository struct {
	transactions map[string]entities.Transaction
	mu           sync.RWMutex
}

func NewInMemoryTransactionRepository() *InMemoryTransactionRepository {
	return &InMemoryTransactionRepository{
		transactions: make(map[string]entities.Transaction),
	}
}

func (r *InMemoryTransactionRepository) Store(t entities.Transaction) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.transactions[t.ID]; exists {
		return errors.New("transaction already exists")
	}

	r.transactions[t.ID] = t
	log.Printf("Transaction with ID %s stored successfully\n", r.transactions[t.ID].ID)
	return nil
}

func (r *InMemoryTransactionRepository) Retrieve(id string) (entities.Transaction, error) {
	log.Println("REPO - Retrieving transaction with ID:", id)
	r.mu.RLock()
	defer r.mu.RUnlock()

	transaction, exists := r.transactions[id]
	if !exists {
		log.Println("Transaction not found with ID:", id)
		return entities.Transaction{}, errors.New("transaction not found")
	}

	return transaction, nil
}
