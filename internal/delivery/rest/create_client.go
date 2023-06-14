package rest

import (
	"net/http"
	"transaction_test_case/internal/domain"

	"github.com/gin-gonic/gin"
)

type CreateClientRequest struct{}

type CreateClient struct {
	Message string
	Client  domain.Client
}

// CreateClient создает нового клиента.
// @Summary Create a new client
// @Description Create a new client
// @ID create-client
// @Accept json
// @Produce json
// @Param client body CreateClientRequest true "Create client"
// @Success 200 {object} CreateClient
// @Router /clients/create [post]
func (h *Handler) CreateClient(c *gin.Context) {
	var req CreateClientRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := h.clientUseCase.CreateClient(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Client created successfully", "client": client})
}
