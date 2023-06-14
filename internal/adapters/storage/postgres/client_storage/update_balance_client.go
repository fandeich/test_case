package clientstorage

import (
	"context"
	"time"
)

const updateBalanceClient = `
UPDATE clients SET balance = $1, updated_at = $2 WHERE id = $3;
`

func (p *Storage) UpdateBalanceClient(ctx context.Context, clientID string, newBalance float64) error {
	currentTime := time.Now()

	_, err := p.db.ExecContext(ctx, updateBalanceClient, newBalance, currentTime, clientID)
	return err
}
