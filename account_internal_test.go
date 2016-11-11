package toshl

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func TestAccountGetQueryString(t *testing.T) {
	a := AccountQueryParams{
		Page:           2,
		PerPage:        1,
		Since:          time.Date(2016, 11, 6, 13, 28, 0, 0, time.Local),
		Status:         "active",
		IncludeDeleted: true,
	}

	assert.Equal(t,
		a.getQueryString(),
		`include_deleted=true&page=2&per_page=1&`+
			`since=2016-11-06T13%3A28%3A00Z&status=active`)
}
