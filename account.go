package toshl

// Account represents a Toshl account
type Account struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Balance        float64  `json:"balance"`
	InitialBalance float64  `json:"initial_balance"`
	Currency       Currency `json:"currency"`
	Median         Median   `json:"median"`
	Status         string   `json:"status"`
	Order          int      `json:"order"`
	Modified       string   `json:"modified"`
	Goal           Goal     `json:"goal"`
}
