package toshl

// Budget represents a Toshl budget
type Budget struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Limit int `json:"limit"`
	Amount float64 `json:"amount"`
	Planned float64 `json:"planned"`
	Median int `json:"median"`
	Currency Currency `json:"currency"`
	From string `json:"from"`
	To string `json:"to"`
	Rollover bool `json:"rollover"`
	Modified string `json:"modified"`
	Recurrence Recurrence `json:"recurrence"`
	Status string `json:"status"`
	Type string `json:"type"`
	Order int `json:"order"`
	Categories []string `json:"categories"`
}
