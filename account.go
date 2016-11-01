package toshl

import "time"

// Account represents a Toshl account
type Account struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Balance        int       `json:"balance"`
	InitialBalance int       `json:"initial_balance"`
	Currency       Currency  `json:"currency"`
	Median         Median    `json:"median"`
	Status         string    `json:"status"`
	Order          int       `json:"order"`
	Modified       time.Time `json:"modified"`
	Goal           Goal      `json:"goal"`
}
