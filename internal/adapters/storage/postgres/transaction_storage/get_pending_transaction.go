package transactionstorage

import (
	"context"
	"transaction_test_case/internal/domain"
)

const getPendingTransactions = `
SELECT id, client_id, amount, status FROM transactions WHERE status = $1 ORDER BY updated_at LIMIT 100;
`

func (s *Storage) GetPendingTransactions(ctx context.Context) ([]*domain.Transaction, error) {
	rows, err := s.db.QueryContext(ctx, getPendingTransactions, domain.Pending)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*domain.Transaction
	for rows.Next() {
		t := &domain.Transaction{}
		var status int
		if err := rows.Scan(&t.ID, &t.ClientID, &t.Amount, &status); err != nil {
			return nil, err
		}
		t.Status = domain.TransactionStatus(status)
		transactions = append(transactions, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
