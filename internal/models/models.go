package models

import "time"

type User struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	Phone   string  `json:"phone"`
	VPA     string  `json:"vpa"`
	Balance float64 `json:"balance"`
}

type Transaction struct {
	ID        string    `json:"id"`
	FromVPA   string    `json:"from_vpa"`
	ToVPA     string    `json:"to_vpa"`
	Amount    string    `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
