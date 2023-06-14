package transactionstorage

import (
	"context"
	"time"
	"transaction_test_case/internal/domain"
)

const updateTransactionStatus = `
UPDATE transactions SET status = $1, updated_at = $2 WHERE id = $3;
`

func (s *Storage) UpdateTransactionStatus(ctx context.Context, transactionID string, status domain.TransactionStatus) error {
	currentTime := time.Now()

	_, err := s.db.ExecContext(ctx, updateTransactionStatus, status, currentTime, transactionID)
	return err
}
