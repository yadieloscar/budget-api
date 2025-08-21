package handler

import (
	"github.com/gin-gonic/gin"
)

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
