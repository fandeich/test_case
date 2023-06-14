package transactionstorage

import (
	"context"
	"database/sql"
	"fmt"
	"transaction_test_case/internal/domain"
)

const getTransactionByID = `
SELECT id, client_id, amount, status, note, created_at, updated_at FROM transactions WHERE id = $1;
`

func (s *Storage) GetTransactionByID(ctx context.Context, transactionID string) (*domain.Transaction, error) {
	row := s.db.QueryRowContext(ctx, getTransactionByID, transactionID)
	transaction := &domain.Transaction{}

	if err := row.Scan(
		&transaction.ID,
		&transaction.ClientID,
		&transaction.Amount,
		&transaction.Status,
		&transaction.Note,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("transaction not found")
		}

		return nil, err
	}

	return transaction, nil
}
