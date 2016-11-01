package toshl

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountDecode(t *testing.T) {
	var account Account
	accountJSON := []byte(`{
        "id": "42",
        "name": "Tesla model S",
        "balance": 3000,
        "initial_balance": 3000,
        "currency": {
            "code": "USD",
            "rate": 1,
            "fixed": false
        },
        "median": {
            "expenses": 55,
            "incomes": 1300
        },
        "status": "active",
        "order": 0,
        "modified": "2012-09-04T13:55:15Z",
        "goal": {
            "amount": 63570,
            "start": "2013-07-01",
            "end": "2015-07-01"
        }
    }`)

	err := json.Unmarshal(accountJSON, &account)

	if err != nil {
		t.Errorf("Cannot decode Account JSON")
	}

	assert.Nil(t, err)
}

func TestAccountDecodeError(t *testing.T) {
	var account Account
	accountJSON := []byte(`{
        "id": "42",
        "name": "Tesla model S",
        "balance": 3000,
        "initial_balance": 3000,
        "currency": {
            "code": "USD",
            "rate": 1,
            "fixed": false
        },
        "median": {
            "expenses": 55,
            "incomes": 1300
        },
        "status": "active",
        "order": 0,
        "modified": "2012-09-04T13:55:15Z",
        "goal": {
            "amount": "abcd",
            "start": "2013-07-01",
            "end": "2015-07-01"
        }
    }`)

	err := json.Unmarshal(accountJSON, &account)
	assert.NotNil(t, err)
}
