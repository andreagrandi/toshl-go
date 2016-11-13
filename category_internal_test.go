package toshl

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func TestCategoryGetQueryString(t *testing.T) {
	c := CategoryQueryParams{
		Page:           2,
		PerPage:        1,
		Since:          time.Date(2016, 11, 6, 13, 28, 0, 0, time.Local),
		Type:           "active",
		Search:         "search_term",
		IncludeDeleted: true,
	}

	assert.Equal(t,
		c.getQueryString(),
		`include_deleted=true&page=2&per_page=1&search=search_term&`+
			`since=2016-11-06T13%3A28%3A00Z&type=active`)
}
