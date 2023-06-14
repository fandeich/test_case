package transactionstorage

import (
	"context"
	"fmt"
	"transaction_test_case/internal/domain"
)

const insertTransaction = `
INSERT INTO transactions (id, client_id, amount, status, note, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7);
`

func (p *Storage) InsertTransaction(ctx context.Context, transaction *domain.Transaction) error {
	_, err := p.db.ExecContext(ctx, insertTransaction,
		transaction.ID,
		transaction.ClientID,
		transaction.Amount,
		transaction.Status,
		transaction.Note,
		transaction.CreatedAt,
		transaction.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to insert transaction: %w", err)
	}
	return nil
}
