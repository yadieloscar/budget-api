// Package service contains business logic for the budget feature.
package service

import (
	"context"
	"errors"

	"github.com/yadieloscar/budget-api/internal/api/budget/model"
	"github.com/yadieloscar/budget-api/internal/api/budget/repo"
)

var (
	ErrInvalidAmount = errors.New("Invalid amount provided. Total amount must be greater than zero")
)

// BudgetService defines business operations for budgets.
type BudgetService interface {
	CreateBudget(ctx context.Context, budgetInput model.CreateBudgetRequest) (*model.Budget, error)
	GetBudget(ctx context.Context, budgetID string) (*model.Budget, error)
}

type svc struct {
	repo repo.BudgetRepository
}

// NewBudgetService creates a new instance of BudgetService.
func NewBudgetService(r repo.BudgetRepository) BudgetService {
	return &svc{repo: r}
}

func (s *svc) CreateBudget(ctx context.Context, budgetInput model.CreateBudgetRequest) (*model.Budget, error) {

	if budgetInput.TotalAmountCents <= 0 {
		return nil, ErrInvalidAmount
	}
	created, err := s.repo.CreateBudget(ctx, budgetInput)
	if err != nil {
		return nil, err
	}
	return &model.Budget{
		ID:               created.ID,
		UserID:           created.UserID,
		Name:             created.Name,
		Month:            created.Month,
		Status:           created.Status,
		AvailableCents:   created.AvailableCents,
		TotalAmountCents: created.TotalAmountCents,
		Currency:         created.Currency,
		CreatedAt:        created.CreatedAt,
		UpdatedAt:        created.UpdatedAt,
	}, nil

}

func (s *svc) GetBudget(ctx context.Context, budgetID string) (*model.Budget, error) {
	budget, err := s.repo.GetBudgetByID(ctx, budgetID)
	if err != nil {
		return nil, err
	}
	return &model.Budget{
		ID:               budget.ID,
		UserID:           budget.UserID,
		Name:             budget.Name,
		Month:            budget.Month,
		Status:           budget.Status,
		AvailableCents:   budget.AvailableCents,
		TotalAmountCents: budget.TotalAmountCents,
		Currency:         budget.Currency,
		CreatedAt:        budget.CreatedAt,
		UpdatedAt:        budget.UpdatedAt,
	}, nil
}
