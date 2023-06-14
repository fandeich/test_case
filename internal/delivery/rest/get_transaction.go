package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type GetTransactionResponse struct {
	TransactionID string    `json:"transaction_id"`
	ClientID      string    `json:"client_id"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"`
	Note          string    `json:"note"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// GetTransaction получает статус транзакции по ее идентификатору.
// @Summary Get a transaction's status
// @Description Get a transaction's status by transaction id
// @ID get-transaction
// @Accept json
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} GetTransactionResponse
// @Router /transactions/{id} [post]
func (h *Handler) GetTransaction(c *gin.Context) {
	transactionID := c.Param("id")

	transaction, err := h.transactionUseCase.GetTransactionByID(c.Request.Context(), transactionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, GetTransactionResponse{
		TransactionID: transaction.ID,
		ClientID:      transaction.ClientID,
		Amount:        transaction.Amount,
		Status:        transaction.Status.String(),
		Note:          transaction.Note,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	})
}
