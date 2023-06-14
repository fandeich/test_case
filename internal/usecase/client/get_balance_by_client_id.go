package client

import "context"

func (uc *Usecase) GetBalanceByClientID(ctx context.Context, clientID string) (float64, error) {
	return uc.clientStorage.GetBalanceByClientID(ctx, clientID)
}
