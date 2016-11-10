package toshl

import (
	"net/url"
	"strconv"
	"time"
)

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

// AccountQueryParams represents a struct of parameters usable
// to List Accounts
type AccountQueryParams struct {
	Page           int
	PerPage        int
	Since          time.Time
	Status         string
	IncludeDeleted bool
}

func (a *AccountQueryParams) getQueryString() string {
	v := url.Values{}

	if a.Page > 0 {
		v.Set("page", strconv.Itoa(a.Page))
	}

	if a.PerPage > 0 {
		v.Set("per_page", strconv.Itoa(a.PerPage))
	}

	if !a.Since.IsZero() {
		v.Set("since", a.Since.Format("2006-01-02T15:04:05Z"))
	}

	if a.Status != "" {
		v.Set("status", a.Status)
	}

	if a.IncludeDeleted {
		v.Set("include_deleted", strconv.FormatBool(a.IncludeDeleted))
	}

	return v.Encode()
}

// AccountsOrderParams describes the order we want for the accounts
type AccountsOrderParams struct {
	Order []string `json:"order"`
}

// AccountsMergeParams describes how we want to merge the accounts
type AccountsMergeParams struct {
	Accounts []string `json:"accounts"`
	Account  string   `json:"account"`
}
