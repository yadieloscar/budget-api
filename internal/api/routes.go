// Package api exposes the top-level HTTP router setup for the service.
// It mounts versioned route groups and registers feature modules.
package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	budget "github.com/yadieloscar/budget-api/internal/api/budget/module"
	"github.com/yadieloscar/budget-api/internal/api/shared/middleware"
)

// SetupRouter constructs the Gin engine, attaches middleware, and registers
// all feature modules under the versioned route group (e.g. /api/v1).
// The provided sql.DB is passed down to modules that require persistence.
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
