package rest

import (
	"transaction_test_case/internal/usecase/interfaces"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	clientUseCase      interfaces.ClientUseCase
	transactionUseCase interfaces.TransactionUseCase
}

func NewHandler(clientUseCase interfaces.ClientUseCase, transactionUseCase interfaces.TransactionUseCase, router *gin.Engine) *Handler {
	h := &Handler{
		clientUseCase:      clientUseCase,
		transactionUseCase: transactionUseCase,
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/clients/create", h.CreateClient)
	router.POST("/clients/:id/balance", h.GetClientBalance)
	router.POST("/transactions/create", h.CreateTransaction)
	router.POST("/transactions/:id", h.GetTransaction)

	return h
}
