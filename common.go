package toshl

// Currency represents a Toshl supported currency
type Currency struct {
	Code  string `json:"code"`
	Rate  int    `json:"rate"`
	Fixed bool   `json:"fixed"`
}

// Median represents a Toshl median
type Median struct {
	Expenses int `json:"expenses"`
	Incomes  int `json:"incomes"`
}

// Goal represents a Toshl goal
type Goal struct {
	Amount int    `json:"amount"`
	Start  string `json:"start"`
	End    string `json:"end"`
}
