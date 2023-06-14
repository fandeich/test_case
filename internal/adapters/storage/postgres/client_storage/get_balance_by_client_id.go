package clientstorage

import (
	"context"
	"database/sql"
	"fmt"
)

const getBalanceByClientID = `
SELECT balance FROM clients WHERE id = $1;
`

func (p *Storage) GetBalanceByClientID(ctx context.Context, clientID string) (float64, error) {
	var balance float64
	rows := p.db.QueryRowContext(ctx, getBalanceByClientID, clientID)

	if err := rows.Scan(&balance); err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("client not found")
		}

		return 0, err
	}

	return balance, nil
}
