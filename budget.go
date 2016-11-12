package toshl

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

// Budget represents a Toshl budget
type Budget struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	Limit      int        `json:"limit"`
	Amount     float64    `json:"amount"`
	Planned    float64    `json:"planned"`
	Median     int        `json:"median"`
	Currency   Currency   `json:"currency"`
	From       string     `json:"from"`
	To         string     `json:"to"`
	Rollover   bool       `json:"rollover"`
	Modified   string     `json:"modified"`
	Recurrence Recurrence `json:"recurrence"`
	Status     string     `json:"status"`
	Type       string     `json:"type"`
	Order      int        `json:"order"`
	Categories []string   `json:"categories"`
}

// BudgetQueryParams represents a struct of parameters usable
// to List Budgets
type BudgetQueryParams struct {
	Page             int
	PerPage          int
	Since            time.Time
	From             time.Time
	To               time.Time
	Tags             []string
	Categories       []string
	Accounts         []string
	Search           string
	IncludeDeleted   bool
	Expand           bool
	HasProblem       bool
	OneIterationOnly bool
	Parent           string
}

func (b *BudgetQueryParams) getQueryString() string {
	v := url.Values{}

	if b.Page > 0 {
		v.Set("page", strconv.Itoa(b.Page))
	}

	if b.PerPage > 0 {
		v.Set("per_page", strconv.Itoa(b.PerPage))
	}

	if !b.Since.IsZero() {
		v.Set("since", b.Since.Format("2006-01-02T15:04:05Z"))
	}

	if !b.From.IsZero() {
		v.Set("from", b.Since.Format("2006-01-02"))
	}

	if !b.To.IsZero() {
		v.Set("to", b.Since.Format("2006-01-02"))
	}

	if len(b.Tags) > 0 {
		v.Set("tags", strings.Join(b.Tags, ","))
	}

	if len(b.Categories) > 0 {
		v.Set("categories", strings.Join(b.Categories, ","))
	}

	if len(b.Accounts) > 0 {
		v.Set("accounts", strings.Join(b.Accounts, ","))
	}

	if b.Search != "" {
		v.Set("search", b.Search)
	}

	if b.IncludeDeleted {
		v.Set("include_deleted", strconv.FormatBool(b.IncludeDeleted))
	}

	if b.Expand {
		v.Set("expand", strconv.FormatBool(b.Expand))
	}

	if b.HasProblem {
		v.Set("has_problem", strconv.FormatBool(b.HasProblem))
	}

	if b.OneIterationOnly {
		v.Set("one_iteration_only", strconv.FormatBool(b.OneIterationOnly))
	}

	if b.Parent != "" {
		v.Set("parent", b.Parent)
	}

	return v.Encode()
}
