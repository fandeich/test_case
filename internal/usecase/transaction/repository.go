package transaction

import (
	"transaction_test_case/internal/usecase/interfaces"
)

type Usecase struct {
	transactionStorage interfaces.TransactionStorage
	clientUseCase      interfaces.ClientUseCase
}

func NewUsecase(transactionStorage interfaces.TransactionStorage, clientUseCase interfaces.ClientUseCase) *Usecase {
	return &Usecase{
		transactionStorage: transactionStorage,
		clientUseCase:      clientUseCase,
	}
}
