// Package handlers contains HTTP handlers for the budget API.
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yadieloscar/budget-api/internal/models"
	"github.com/yadieloscar/budget-api/internal/services"
)

// BudgetHandler handles budget-related HTTP requests.
type BudgetHandler struct {
	svc services.BudgetService
}

// NewBudgetHandler returns a handler bound to a BudgetService.
func NewBudgetHandler(svc services.BudgetService) *BudgetHandler {
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

	var budgetInput models.CreateBudgetRequest
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
		if err == services.ErrInvalidAmount {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdBudget)
}

// RegisterBudgetRoutes mounts the budget endpoints under the provided router group.
func RegisterBudgetRoutes(r *gin.RouterGroup, h *BudgetHandler) {
	api := r.Group("/budget")
	{
		api.GET("/", h.GetBudget)
		api.POST("/create", h.CreateBudget)
	}
}
