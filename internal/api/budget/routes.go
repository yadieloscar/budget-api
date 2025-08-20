package budget

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	api := r.Group("/budget")
	{
		api.GET("/", GetBudget)
		// api.POST("/items", handlers.CreateItem)
		// api.GET("/items/:id", handlers.GetItem)
		// api.PUT("/items/:id", handlers.UpdateItem)
		// api.DELETE("/items/:id", handlers.DeleteItem)
	}
}
