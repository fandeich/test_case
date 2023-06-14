package transaction

import (
	"context"
	"fmt"
	"transaction_test_case/internal/domain"
)

func (uc *Usecase) ProcessPendingTransactions(ctx context.Context) error {
	transactions, err := uc.transactionStorage.GetPendingTransactions(ctx)
	if err != nil {
		return fmt.Errorf("error getting pending transactions: %w", err)
	}

	for _, transaction := range transactions {
		balance, err := uc.clientUseCase.GetBalanceByClientID(ctx, transaction.ClientID)
		if err != nil {
			return fmt.Errorf("error getting balance for client %s: %w", transaction.ClientID, err)
		}

		newBalance := balance + transaction.Amount

		if newBalance < 0 {
			if err := uc.transactionStorage.UpdateTransactionStatus(ctx, transaction.ID, domain.Rejected); err != nil {
				return fmt.Errorf("error updating transaction status: %w", err)
			}
			continue
		}

		if err := uc.transactionStorage.UpdateTransactionAndBalance(ctx, transaction.ID, domain.Completed, transaction.ClientID, newBalance); err != nil {
			return fmt.Errorf("error updating transaction status and balance: %w", err)
		}
	}

	return nil
}
