package transaction

import (
	"context"
	"transaction_test_case/internal/domain"
)

func (uc *Usecase) GetTransactionByID(ctx context.Context, transactionID string) (*domain.Transaction, error) {
	return uc.transactionStorage.GetTransactionByID(ctx, transactionID)
}
