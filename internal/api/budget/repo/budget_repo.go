package repo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/yadieloscar/budget-api/internal/api/budget/model"
)

var (
	ErrBudgetNotFound = errors.New("Budget not found")
)

type BudgetRepository interface {
	// CreateBudget creates a new budget entry in the repository
	CreateBudget(ctx context.Context, budgetInput model.CreateBudgetRequest) (model.Budget, error)
	GetBudgetByID(ctx context.Context, id string) (model.Budget, error)
}

type postgresBudgetRepo struct {
	db *sql.DB // Assuming you have a database connection
}

func NewBudgetRepo(db *sql.DB) BudgetRepository {
	return &postgresBudgetRepo{db: db}
}

func (r *postgresBudgetRepo) CreateBudget(ctx context.Context, budgetInput model.CreateBudgetRequest) (model.Budget, error) {
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
		string(model.StatusActive),   // Convert BudgetStatus to string for storage
		budgetInput.TotalAmountCents, // Initially, available amount equals total amount
	).Scan(&id, &status, &createdAt, &updatedAt)

	if err != nil {
		return model.Budget{}, err
	}
	return model.Budget{
		ID:               id,
		UserID:           budgetInput.UserID,
		Name:             budgetInput.Name,
		Month:            budgetInput.Month,
		Status:           model.BudgetStatus(status), // Convert string back to BudgetStatus
		AvailableCents:   budgetInput.TotalAmountCents,
		TotalAmountCents: budgetInput.TotalAmountCents,
		Currency:         budgetInput.Currency,
		CreatedAt:        createdAt, // You can set this to the actual created_at timestamp if
		UpdatedAt:        updatedAt, // Same as above, or you can set it to the
	}, nil
}

func (r *postgresBudgetRepo) GetBudgetByID(ctx context.Context, id string) (model.Budget, error) {
	query := `
		SELECT id, user_id, name, month, status, available_cents, total_amount_cents, currency, created_at, updated_at
		FROM budgets WHERE id = $1`

	var budget model.Budget
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
			return model.Budget{}, ErrBudgetNotFound
		}
		return model.Budget{}, err
	}
	budget.Status = model.BudgetStatus(budget.Status) // Convert string back to BudgetStatus
	return budget, nil
}
