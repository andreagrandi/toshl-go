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

// Recurrence represents a Toshl recurrence
type Recurrence struct {
	Frequency string `json:"frequency"`
	Interval  int    `json:"interval"`
	Start     string `json:"start"`
	Iteration int    `json:"iteration"`
}

// CategoryCounts represents a Toshl count
type CategoryCounts struct {
	Entries int `json:"entries"`
	Tags    int `json:"tags"`
}
