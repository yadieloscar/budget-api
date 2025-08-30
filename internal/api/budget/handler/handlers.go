// Package handler contains HTTP handlers for the budget feature.
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yadieloscar/budget-api/internal/api/budget/model"
	"github.com/yadieloscar/budget-api/internal/api/budget/service"
)

type BudgetHandler struct {
	svc service.BudgetService
}

// NewBudgetHandler returns a handler bound to a BudgetService.
func NewBudgetHandler(svc service.BudgetService) *BudgetHandler {
	return &BudgetHandler{svc: svc}
}

// GetBudget handles GET requests for retrieving budget information
func (h *BudgetHandler) GetBudget(c *gin.Context) {
	uid, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}
	budget, err := h.svc.GetBudget(c.Request.Context(), uid.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if budget == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Budget not found"})
		return
	}
	c.JSON(http.StatusOK, budget)
}

// CreateBudget handles POST requests for creating a new budget
func (h *BudgetHandler) CreateBudget(c *gin.Context) {

	var budgetInput model.CreateBudgetRequest
	if err := c.ShouldBindJSON(&budgetInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	uid, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}
	budgetInput.UserID = uid.(string)
	createdBudget, err := h.svc.CreateBudget(c.Request.Context(), budgetInput)
	if err != nil {
		if err == service.ErrInvalidAmount {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdBudget)
}
