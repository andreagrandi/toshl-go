package toshl

import (
	"testing"

	"time"

	"github.com/stretchr/testify/assert"
)

func TestBudgetGetQueryString(t *testing.T) {
	b := BudgetQueryParams{
		Page:             2,
		PerPage:          1,
		Since:            time.Date(2016, 11, 6, 13, 28, 0, 0, time.Local),
		From:             time.Date(2016, 11, 6, 13, 28, 0, 0, time.Local),
		To:               time.Date(2016, 11, 6, 13, 28, 0, 0, time.Local),
		Tags:             []string{"tag1", "tag2"},
		Categories:       []string{"cat1", "cat2"},
		Accounts:         []string{"id1", "id2"},
		Search:           "search_term",
		IncludeDeleted:   true,
		Expand:           true,
		HasProblem:       true,
		OneIterationOnly: true,
		Parent:           "p1",
	}

	assert.Equal(t,
		b.getQueryString(),
		`accounts=id1%2Cid2&categories=cat1%2Ccat2&expand=true&`+
			`from=2016-11-06&has_problem=true&include_deleted=true&`+
			`one_iteration_only=true&page=2&parent=p1&per_page=1&`+
			`search=search_term&since=2016-11-06T13%3A28%3A00Z&`+
			`tags=tag1%2Ctag2&to=2016-11-06`)
}
