package interfaces

import (
	"context"
	"transaction_test_case/internal/domain"
)

type TransactionUseCase interface {
	GetTransactionByID(ctx context.Context, transactionID string) (*domain.Transaction, error)
	CreateTransaction(ctx context.Context, clientID string, amount float64) (*domain.Transaction, error)
	ProcessPendingTransactions(ctx context.Context) error
}

type ClientUseCase interface {
	CreateClient(ctx context.Context) (*domain.Client, error)
	GetBalanceByClientID(ctx context.Context, clientID string) (float64, error)
	UpdateBalanceClient(ctx context.Context, clientID string, balance float64) error
}
