// Package repo provides persistence for entities.
package repo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/yadieloscar/budget-api/internal/models"
)

var (
	ErrBudgetNotFound = errors.New("Budget not found")
)

// BudgetRepository defines persistence operations for budgets.
type BudgetRepository interface {
	// CreateBudget creates a new budget entry in the repository
	CreateBudget(ctx context.Context, budgetInput models.CreateBudgetRequest) (models.Budget, error)
	GetBudgetByID(ctx context.Context, id string) (models.Budget, error)
}

// postgresBudgetRepo implements BudgetRepository using a PostgreSQL database.
type postgresBudgetRepo struct {
	db *sql.DB // Assuming you have a database connection
}

// NewBudgetRepo constructs a PostgreSQL-backed budget repository.
func NewBudgetRepo(db *sql.DB) BudgetRepository {
	return &postgresBudgetRepo{db: db}
}

// CreateBudget inserts a budget row and returns the created entity.
func (r *postgresBudgetRepo) CreateBudget(ctx context.Context, budgetInput models.CreateBudgetRequest) (models.Budget, error) {
	query := `
        INSERT INTO budgets (
            user_id, 
            name,
            total_amount_cents, 
            currency, 
            month,
            status,  -- Status will always be 'active' for new budgets
            available_cents,
            created_at,
            updated_at
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, NOW(), NOW()
        ) RETURNING id, status, created_at, updated_at`

	var id string
	var status string
	var createdAt, updatedAt time.Time
	err := r.db.QueryRowContext(
		ctx,
		query,
		budgetInput.UserID,
		budgetInput.Name,
		budgetInput.TotalAmountCents,
		budgetInput.Currency,
		budgetInput.Month,
		string(models.StatusActive),  // Convert BudgetStatus to string for storage
		budgetInput.TotalAmountCents, // Initially, available amount equals total amount
	).Scan(&id, &status, &createdAt, &updatedAt)

	if err != nil {
		return models.Budget{}, err
	}
	return models.Budget{
		ID:               id,
		UserID:           budgetInput.UserID,
		Name:             budgetInput.Name,
		Month:            budgetInput.Month,
		Status:           models.BudgetStatus(status), // Convert string back to BudgetStatus
		AvailableCents:   budgetInput.TotalAmountCents,
		TotalAmountCents: budgetInput.TotalAmountCents,
		Currency:         budgetInput.Currency,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
	}, nil
}

// GetBudgetByID loads a budget by its identifier.
func (r *postgresBudgetRepo) GetBudgetByID(ctx context.Context, id string) (models.Budget, error) {
	query := `
		SELECT id, user_id, name, month, status, available_cents, total_amount_cents, currency, created_at, updated_at
		FROM budgets WHERE id = $1`

	var budget models.Budget
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&budget.ID,
		&budget.UserID,
		&budget.Name,
		&budget.Month,
		&budget.Status,
		&budget.AvailableCents,
		&budget.TotalAmountCents,
		&budget.Currency,
		&budget.CreatedAt,
		&budget.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Budget{}, ErrBudgetNotFound
		}
		return models.Budget{}, err
	}
	budget.Status = models.BudgetStatus(budget.Status) // Convert string back to BudgetStatus
	return budget, nil
}
