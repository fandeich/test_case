package client

import (
	"transaction_test_case/internal/usecase/interfaces"
)

type Usecase struct {
	clientStorage interfaces.ClientStorage
}

func NewUsecase(clientStorage interfaces.ClientStorage) *Usecase {
	return &Usecase{
		clientStorage: clientStorage,
	}
}
