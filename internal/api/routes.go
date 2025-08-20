package api

import (
	"github.com/gin-gonic/gin"
	"github.com/yadieloscar/budget-api/internal/api/budget"
	"github.com/yadieloscar/budget-api/internal/api/categories"
	"github.com/yadieloscar/budget-api/internal/api/savings"
	"github.com/yadieloscar/budget-api/internal/api/shared/middleware"
	"github.com/yadieloscar/budget-api/internal/api/transactions"
	"github.com/yadieloscar/budget-api/internal/api/users"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())

	// API v1 group
	v1 := r.Group("/api/v1")
	{
		// Mount all feature routes
		budget.RegisterRoutes(v1)
		categories.RegisterRoutes(v1)
		savings.RegisterRoutes(v1)
		transactions.RegisterRoutes(v1)
		users.RegisterRoutes(v1)
	}

	return r
}
