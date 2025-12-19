// Package handlers contains HTTP handlers and routing for the budget API.
package handlers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/yadieloscar/budget-api/internal/middleware"
	"github.com/yadieloscar/budget-api/internal/repo"
	"github.com/yadieloscar/budget-api/internal/services"
)

// SetupRouter constructs the Gin engine, attaches middleware, and registers
// all handlers under the versioned route group (e.g. /api/v1).
// The provided sql.DB is passed down to repositories that require persistence.
func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.CORSMiddleware())

	v1 := r.Group("/api/v1")

	// Wire up budget feature: repo -> service -> handler
	budgetRepo := repo.NewBudgetRepo(db)
	budgetService := services.NewBudgetService(budgetRepo)
	budgetHandler := NewBudgetHandler(budgetService)
	RegisterBudgetRoutes(v1, budgetHandler)

	// Wire up placeholder handlers
	categoriesHandler := NewCategoriesHandler()
	RegisterCategoriesRoutes(v1, categoriesHandler)

	savingsHandler := NewSavingsHandler()
	RegisterSavingsRoutes(v1, savingsHandler)

	transactionsHandler := NewTransactionsHandler()
	RegisterTransactionsRoutes(v1, transactionsHandler)

	usersHandler := NewUsersHandler()
	RegisterUsersRoutes(v1, usersHandler)

	return r
}
