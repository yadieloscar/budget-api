package models

type Budget struct {
    ID          string  `json:"id"`
    Name        string  `json:"name"`
    Amount      float64 `json:"amount"`
    Description string  `json:"description"`
}