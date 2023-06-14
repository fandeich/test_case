package transactionstorage

import (
	"context"
	"fmt"
	"time"
	"transaction_test_case/internal/domain"
)

const updateBalanceClient = `
UPDATE clients SET balance = $1, updated_at = $2 WHERE id = $3;
`

func (s *Storage) UpdateTransactionAndBalance(ctx context.Context, transactionID string, status domain.TransactionStatus, clientID string, newBalance float64) error {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error beginning transaction: %w", err)
	}

	currentTime := time.Now()

	_, err = tx.ExecContext(ctx, updateTransactionStatus, status, currentTime, transactionID)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("error updating transaction status: %w", err)
	}

	_, err = tx.ExecContext(ctx, updateBalanceClient, newBalance, currentTime, clientID)
	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("error updating client balance: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}
