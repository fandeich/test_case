package client

import "context"

func (uc Usecase) UpdateBalanceClient(ctx context.Context, clientID string, balance float64) error {
	return uc.clientStorage.UpdateBalanceClient(ctx, clientID, balance)
}
