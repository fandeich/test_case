package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateTransactionRequest struct {
	ClientID string  `json:"client_id"`
	Amount   float64 `json:"amount"`
}

type CreateTransactionResponse struct {
	TransactionID string    `json:"transaction_id"`
	ClientID      string    `json:"client_id"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`
	Note          string    `json:"note"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// CreateTransaction создает новую транзакцию для указанного клиента с указанной суммой.
// @Summary Create a new transaction
// @Description Create a new transaction for a specific client with an amount
// @ID create-transaction
// @Accept json
// @Produce json
// @Param transaction body CreateTransactionRequest true "Create transaction"
// @Success 200 {object} CreateTransactionResponse
// @Router /transactions/create [post]
func (h *Handler) CreateTransaction(c *gin.Context) {
	var request CreateTransactionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := h.transactionUseCase.CreateTransaction(c.Request.Context(), request.ClientID, request.Amount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, CreateTransactionResponse{
		TransactionID: transaction.ID,
		ClientID:      transaction.ClientID,
		Amount:        transaction.Amount,
		Status:        transaction.Status.String(),
		Note:          transaction.Note,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	})
}
