package client

import (
	"context"
	"transaction_test_case/internal/domain"
)

func (uc *Usecase) CreateClient(ctx context.Context) (*domain.Client, error) {
	return uc.clientStorage.CreateClient(ctx)
}
