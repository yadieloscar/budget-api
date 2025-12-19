// Package handlers contains HTTP handlers for the budget API.
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TransactionsHandler handles transaction-related HTTP requests.
type TransactionsHandler struct{}

// NewTransactionsHandler returns a new transactions handler.
func NewTransactionsHandler() *TransactionsHandler {
	return &TransactionsHandler{}
}

// GetTransactions handles GET requests for retrieving transactions
func (h *TransactionsHandler) GetTransactions(c *gin.Context) {
	// Implementation will go here
	c.JSON(http.StatusOK, gin.H{"message": "GetTransactions called"})
}

// CreateTransaction handles POST requests for creating a new transaction
func (h *TransactionsHandler) CreateTransaction(c *gin.Context) {
	// Implementation will go here
	c.JSON(http.StatusCreated, gin.H{"message": "CreateTransaction called"})
}

// UpdateTransaction handles PUT requests for updating an existing transaction
func (h *TransactionsHandler) UpdateTransaction(c *gin.Context) {
	// Implementation will go here
	c.JSON(http.StatusOK, gin.H{"message": "UpdateTransaction called"})
}

// DeleteTransaction handles DELETE requests for removing a transaction
func (h *TransactionsHandler) DeleteTransaction(c *gin.Context) {
	// Implementation will go here
	c.JSON(http.StatusNoContent, nil)
}

// RegisterTransactionsRoutes mounts the transactions endpoints under the provided router group.
func RegisterTransactionsRoutes(r *gin.RouterGroup, h *TransactionsHandler) {
	api := r.Group("/transactions")
	{
		api.GET("/", h.GetTransactions)
		api.POST("/", h.CreateTransaction)
		api.PUT("/:id", h.UpdateTransaction)
		api.DELETE("/:id", h.DeleteTransaction)
	}
}
