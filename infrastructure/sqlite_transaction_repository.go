package infrastructure

import (
	"database/sql"
	"errors"
	"log"
	"transaction/domain/entities"

	_ "github.com/mattn/go-sqlite3"
)

type SqliteTransactionRepository struct {
	db *sql.DB
}

func NewSqliteTransactionRepository(db *sql.DB) *SqliteTransactionRepository {
	return &SqliteTransactionRepository{db: db}
}

func (r *SqliteTransactionRepository) Store(t entities.Transaction) error {
	tx, err := r.db.Begin()
	log.Printf("Starting transaction storage - ID: %s, Description: %s, Transaction Date: %s, Amount: %f", t.ID, t.Description, t.TransactionDate, t.Amount)
	if err != nil {
		log.Println("Error beginning transaction:", err)
		return err
	}
	stmt, err := tx.Prepare("INSERT INTO transactions(id, description, transaction_date, amount) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Println("Error preparing statement:", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(t.ID, t.Description, t.TransactionDate, t.Amount)
	if err != nil {
		tx.Rollback()
		log.Println("Error executing statement and rolling back transaction:", err)
		return err
	}
	tx.Commit()
	log.Println("Transaction stored successfully")
	return nil
}

func (r *SqliteTransactionRepository) Retrieve(id string) (entities.Transaction, error) {
	var transaction entities.Transaction
	err := r.db.QueryRow("SELECT id, description, transaction_date, amount FROM transactions WHERE id = ?", id).Scan(&transaction.ID, &transaction.Description, &transaction.TransactionDate, &transaction.Amount)
	if err != nil {
		log.Println("Transaction not found with ID:", id)
		return entities.Transaction{}, errors.New("transaction not found")
	}
	log.Printf("Retrieved transaction - ID: %s, Description: %s, Amount: %f", transaction.ID, transaction.Description, transaction.Amount)
	return transaction, nil
}

func InitDB(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS transactions (id TEXT PRIMARY KEY, description TEXT, transaction_date DATE, amount REAL)")
	if err != nil {
		return nil, err
	}

	log.Println("Database initialized successfully")
	return db, nil
}
