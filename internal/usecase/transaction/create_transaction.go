package transaction

import (
	"context"
	"fmt"
	"time"
	"transaction_test_case/internal/domain"

	"github.com/google/uuid"
)

func (uc *Usecase) CreateTransaction(ctx context.Context, clientID string, amount float64) (*domain.Transaction, error) {
	balance, err := uc.clientUseCase.GetBalanceByClientID(ctx, clientID)
	if err != nil {
		return nil, fmt.Errorf("Error with getting client balance: %w", err)
	}

	transaction := &domain.Transaction{
		ID:        uuid.New().String(),
		ClientID:  clientID,
		Amount:    amount,
		Status:    domain.Pending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	totalAmount := balance + amount

	if totalAmount < 0 {
		transaction.Status = domain.Rejected
		transaction.Note = "Insufficient balance"
	}

	err = uc.transactionStorage.InsertTransaction(ctx, transaction)
	if err != nil {
		return nil, fmt.Errorf("Error with getting inserting transaction: %w", err)
	}

	return transaction, nil
}
