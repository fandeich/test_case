package main

import (
	"context"
	"os"
	_ "transaction_test_case/docs"
	"transaction_test_case/internal/adapters/storage/postgres"
	clientstorage "transaction_test_case/internal/adapters/storage/postgres/client_storage"
	transactionstorage "transaction_test_case/internal/adapters/storage/postgres/transaction_storage"
	"transaction_test_case/internal/delivery/rest"
	"transaction_test_case/internal/usecase/client"
	"transaction_test_case/internal/usecase/transaction"
	"transaction_test_case/internal/worker"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	//os.Setenv("DB_CONN", "postgresql://postgres:postgres@localhost:5432/transaction_test_case?sslmode=disable")
	connString := os.Getenv("DB_CONN")

	postgresDB := postgres.InitPostgres(ctx, connString)
	defer postgresDB.Close()

	transactionStorage := transactionstorage.NewStorage(postgresDB)
	clientStorage := clientstorage.NewStorage(postgresDB)

	clientUseCase := client.NewUsecase(clientStorage)
	transactionUsecase := transaction.NewUsecase(transactionStorage, clientUseCase)

	transactionWorker := worker.NewTransactionWorker(transactionUsecase)
	go transactionWorker.Start(ctx)

	router := gin.Default()

	// setup routes
	rest.NewHandler(clientUseCase, transactionUsecase, router)

	// start server
	router.Run(":8080")
}
