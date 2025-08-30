// Package module wires the budget feature's repository, service, and handlers
// into the HTTP router.
package module

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/yadieloscar/budget-api/internal/api/budget/handler"
	"github.com/yadieloscar/budget-api/internal/api/budget/repo"
	"github.com/yadieloscar/budget-api/internal/api/budget/service"
)

type Module struct {
	db *sql.DB
}

// New constructs a budget module with the provided database handle.
func New(db *sql.DB) *Module {
	return &Module{
		db: db,
	}
}

// Register registers the budget routes under the given router group.
func (m *Module) Register(g *gin.RouterGroup) {
	budgetRepo := repo.NewBudgetRepo(m.db)
	budgetService := service.NewBudgetService(budgetRepo)
	budgetHandler := handler.NewBudgetHandler(budgetService)
	handler.RegisterRoutes(g, budgetHandler)

}
