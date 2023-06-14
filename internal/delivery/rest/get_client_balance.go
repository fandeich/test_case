package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetBalanceResponse struct {
	ClientID string  `json:"client_id"`
	Balance  float64 `json:"balance"`
}

// GetClientBalance получает баланс клиента по его идентификатору.
// @Summary Get a client's balance
// @Description Get a client's balance by client's id
// @ID get-balance
// @Accept json
// @Produce json
// @Param id path string true "Client ID"
// @Success 200 {object} GetBalanceResponse
// @Router /clients/{id}/balance [post]
func (h *Handler) GetClientBalance(c *gin.Context) {
	clientID := c.Param("id")

	balance, err := h.clientUseCase.GetBalanceByClientID(c.Request.Context(), clientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := GetBalanceResponse{
		ClientID: clientID,
		Balance:  balance,
	}

	c.JSON(http.StatusOK, response)
}
