package repositories

import "transaction/domain/entities"

type TransactionRepository interface {
	Retrieve(id string) (entities.Transaction, error)
	Store(t entities.Transaction) error
}
