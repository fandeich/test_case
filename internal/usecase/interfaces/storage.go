package interfaces

import (
	"context"
	"transaction_test_case/internal/domain"
)

type TransactionStorage interface {
	GetTransactionByID(ctx context.Context, transactionID string) (*domain.Transaction, error)
	InsertTransaction(ctx context.Context, transaction *domain.Transaction) error
	GetPendingTransactions(ctx context.Context) ([]*domain.Transaction, error)
	UpdateTransactionStatus(ctx context.Context, transactionID string, status domain.TransactionStatus) error
	UpdateTransactionAndBalance(ctx context.Context, transactionID string, status domain.TransactionStatus, clientID string, balance float64) error
}

type ClientStorage interface {
	CreateClient(ctx context.Context) (*domain.Client, error)
	GetBalanceByClientID(ctx context.Context, clientID string) (float64, error)
	UpdateBalanceClient(ctx context.Context, clientID string, balance float64) error
}
