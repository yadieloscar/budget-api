package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	budget "github.com/yadieloscar/budget-api/internal/api/budget/module"
	"github.com/yadieloscar/budget-api/internal/api/shared/middleware"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())

	v1 := r.Group("/api/v1")
	modules := []interface{ Register(g *gin.RouterGroup) }{
		budget.New(db),
	}

	for _, module := range modules {
		module.Register(v1)
	}

	return r
}
