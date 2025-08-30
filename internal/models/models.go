// Package models holds shared or legacy model definitions.
package models

// Budget is a simple example model kept for reference.
type Budget struct {
    ID          string  `json:"id"`
    Name        string  `json:"name"`
    Amount      float64 `json:"amount"`
    Description string  `json:"description"`
}
