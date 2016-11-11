package toshl_test

import (
	"encoding/json"
	"testing"

	toshl "github.com/andreagrandi/toshl-go"
	"github.com/stretchr/testify/assert"
)

func TestBudgetDecode(t *testing.T) {
	var budget toshl.Budget
	budgetJSON := []byte(`{
        "id": "42",
        "name": "Monthly budget",
        "limit": 1000,
        "amount": 78.4,
        "planned": 12.6,
        "median": 22,
        "currency": {
            "code": "USD",
            "rate": 1,
            "fixed": false
        },
        "from": "2013-02-01",
        "to": "2013-02-30",
        "rollover": false,
        "modified": "2013-06-27T14:14:03+00:00Z",
        "recurrence": {
            "frequency": "monthly",
            "interval": 1,
            "start": "2012-06-01",
            "iteration": 4
        },
        "status": "active",
        "type": "regular",
        "order": 0,
        "categories": ["42"]
    }`)

	err := json.Unmarshal(budgetJSON, &budget)

	if err != nil {
		t.Errorf("Cannot decode Budget JSON")
	}

	assert.Nil(t, err)
}
