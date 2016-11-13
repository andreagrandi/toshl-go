package toshl

import (
	"net/url"
	"strconv"
	"time"
)

// Category represents a Toshl category
type Category struct {
	ID       string         `json:"id"`
	Name     string         `json:"name"`
	Modified time.Time      `json:"modified"`
	Type     string         `json:"type"`
	Deleted  bool           `json:"deleted"`
	Counts   CategoryCounts `json:"counts"`
}

// CategoryQueryParams represents a struct of parameters usable
// to List Categories
type CategoryQueryParams struct {
	Page           int
	PerPage        int
	Since          time.Time
	Type           string
	Search         string
	IncludeDeleted bool
}

func (c *CategoryQueryParams) getQueryString() string {
	v := url.Values{}

	if c.Page > 0 {
		v.Set("page", strconv.Itoa(c.Page))
	}

	if c.PerPage > 0 {
		v.Set("per_page", strconv.Itoa(c.PerPage))
	}

	if !c.Since.IsZero() {
		v.Set("since", c.Since.Format("2006-01-02T15:04:05Z"))
	}

	if c.Type != "" {
		v.Set("type", c.Type)
	}

	if c.Search != "" {
		v.Set("search", c.Search)
	}

	if c.IncludeDeleted {
		v.Set("include_deleted", strconv.FormatBool(c.IncludeDeleted))
	}

	return v.Encode()
}
