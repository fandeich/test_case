package clientstorage

import (
	"context"
	"fmt"
	"time"
	"transaction_test_case/internal/domain"

	"github.com/google/uuid"
)

const createClientQuery = `
INSERT INTO clients (id, balance, created_at, updated_at) VALUES ($1, $2, $3, $4);
`

func (s *Storage) CreateClient(ctx context.Context) (*domain.Client, error) {
	newClientID := uuid.New().String()

	currentTime := time.Now()

	_, err := s.db.ExecContext(ctx, createClientQuery, newClientID, 0, currentTime, currentTime)

	if err != nil {
		return nil, fmt.Errorf("cannot create new client: %w", err)
	}

	newClient := &domain.Client{
		ID:        newClientID,
		Balance:   0,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	return newClient, nil
}
