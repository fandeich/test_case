package worker

import (
	"context"
	"fmt"
	"log"
	"time"
	"transaction_test_case/internal/usecase/interfaces"
)

type TransactionWorker struct {
	transactionUseCase interfaces.TransactionUseCase
}

func NewTransactionWorker(transactionUseCase interfaces.TransactionUseCase) *TransactionWorker {
	return &TransactionWorker{
		transactionUseCase: transactionUseCase,
	}
}

func (w *TransactionWorker) Start(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			w.processTransactions(ctx)
		}
	}
}

func (w *TransactionWorker) processTransactions(ctx context.Context) {
	err := w.transactionUseCase.ProcessPendingTransactions(ctx)
	if err != nil {
		log.Println(fmt.Errorf("Error with ProcessPendingTransactions: %w", err))
		return
	}
}
