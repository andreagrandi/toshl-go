package toshl

// Currency represents a Toshl supported currency
type Currency struct {
	Code  string  `json:"code"`
	Rate  float64 `json:"rate"`
	Fixed bool    `json:"fixed"`
}

// Median represents a Toshl median
type Median struct {
	Expenses float64 `json:"expenses"`
	Incomes  float64 `json:"incomes"`
}

// Goal represents a Toshl goal
type Goal struct {
	Amount float64 `json:"amount"`
	Start  string  `json:"start"`
	End    string  `json:"end"`
}
