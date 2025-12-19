// Package services contains business logic for the budget API.
package services

import (
	"context"
	"errors"

	"github.com/yadieloscar/budget-api/internal/models"
	"github.com/yadieloscar/budget-api/internal/repo"
)

var (
	ErrInvalidAmount = errors.New("Invalid amount provided. Total amount must be greater than zero")
)

// BudgetService defines business operations for budgets.
type BudgetService interface {
	CreateBudget(ctx context.Context, budgetInput models.CreateBudgetRequest) (*models.Budget, error)
	GetBudget(ctx context.Context, budgetID string) (*models.Budget, error)
}

type budgetSvc struct {
	repo repo.BudgetRepository
}

// NewBudgetService creates a new instance of BudgetService.
func NewBudgetService(r repo.BudgetRepository) BudgetService {
	return &budgetSvc{repo: r}
}

func (s *budgetSvc) CreateBudget(ctx context.Context, budgetInput models.CreateBudgetRequest) (*models.Budget, error) {

	if budgetInput.TotalAmountCents <= 0 {
		return nil, ErrInvalidAmount
	}
	created, err := s.repo.CreateBudget(ctx, budgetInput)
	if err != nil {
		return nil, err
	}
	return &models.Budget{
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

func (s *budgetSvc) GetBudget(ctx context.Context, budgetID string) (*models.Budget, error) {
	budget, err := s.repo.GetBudgetByID(ctx, budgetID)
	if err != nil {
		return nil, err
	}
	return &models.Budget{
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
