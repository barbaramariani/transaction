package usecases_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"transaction/application/usecases"
	"transaction/domain/entities"
)

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) Retrieve(id string) (entities.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(entities.Transaction), args.Error(1)
}

func (m *MockTransactionRepository) Store(t entities.Transaction) error {
	args := m.Called(t)
	return args.Error(0)
}

type MockExchangeService struct {
	mock.Mock
}

func (m *MockExchangeService) GetRate(date time.Time, currency string) (float64, error) {
	args := m.Called(date, currency)
	return args.Get(0).(float64), args.Error(1)
}

var (
	repo    *MockTransactionRepository
	service *MockExchangeService
	usecase *usecases.RetrieveTransaction
)

func setup() {
	repo = &MockTransactionRepository{}
	service = &MockExchangeService{}
	usecase = &usecases.RetrieveTransaction{
		TransactionRepo: repo,
		ExchangeService: service,
	}
}

func TestRetrieveTransaction_Execute_Success(t *testing.T) {
	setup()
	repo := &MockTransactionRepository{}
	service := &MockExchangeService{}
	usecase := &usecases.RetrieveTransaction{
		TransactionRepo: repo,
		ExchangeService: service,
	}

	transaction := entities.Transaction{
		ID:              "90445774-14b2-44e1-8caf-3afd9bffcb83",
		Description:     "test",
		TransactionDate: time.Now(),
		Amount:          100.0,
	}

	repo.On("Retrieve", "90445774-14b2-44e1-8caf-3afd9bffcb83").Return(transaction, nil)
	service.On("GetRate", mock.Anything, "USD").Return(0.9, nil)

	result, err := usecase.Execute("90445774-14b2-44e1-8caf-3afd9bffcb83", "USD")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	expectedAmount := transaction.Amount * 0.9
	if result.ConvertedAmount != expectedAmount {
		t.Errorf("Expected amount %v, got %v", expectedAmount, result.ConvertedAmount)
	}

	repo.AssertExpectations(t)
	service.AssertExpectations(t)
}

func TestRetrieveTransaction_Execute_ErrorRetrievingTransaction(t *testing.T) {
	setup()
	repo := &MockTransactionRepository{}
	service := &MockExchangeService{}
	usecase := &usecases.RetrieveTransaction{
		TransactionRepo: repo,
		ExchangeService: service,
	}

	repo.On("Retrieve", "90445774-14b2-44e1-8caf-3afd9bffcb83").Return(entities.Transaction{}, errors.New("error"))

	_, err := usecase.Execute("90445774-14b2-44e1-8caf-3afd9bffcb83", "USD")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	repo.AssertExpectations(t)
	service.AssertExpectations(t)
}
