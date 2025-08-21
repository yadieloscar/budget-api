package model

import "time"

type BudgetStatus string

const (
	StatusActive   BudgetStatus = "active"
	StatusInactive BudgetStatus = "inactive"
	StatusArchived BudgetStatus = "archived"
)

type Budget struct {
	ID               string       `json:"id"`
	UserID           string       `json:"user_id"`
	Name             string       `json:"name"`
	Month            string       `json:"month"`
	Status           BudgetStatus `json:"status"`
	AvailableCents   int64        `json:"available_amount"`
	TotalAmountCents int64        `json:"total_amount"`
	Currency         string       `json:"currency"`
	CreatedAt        time.Time    `json:"created_at"`
	UpdatedAt        time.Time    `json:"updated_at"`
}

type CreateBudgetRequest struct {
	UserID           string `json:"user_id" binding:"required"`
	Name             string `json:"name" binding:"required"`
	TotalAmountCents int64  `json:"total_amount" binding:"required"`
	Currency         string `json:"currency" binding:"required"`
	Month            string `json:"month" binding:"required"`
}
