// Package handler contains HTTP handlers and routing for the budget feature.
package handler

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes mounts the budget endpoints under the provided router group.
func RegisterRoutes(r *gin.RouterGroup, h *BudgetHandler) {
	api := r.Group("/budget")
	{
		api.GET("/", h.GetBudget)
		api.POST("/create", h.CreateBudget)
		// api.GET("/items/:id", handlers.GetItem)
		// api.PUT("/items/:id", handlers.UpdateItem)
		// api.DELETE("/items/:id", handlers.DeleteItem)
	}
}
