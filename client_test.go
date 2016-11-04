package toshl_test

import (
	"testing"

	toshl "github.com/andreagrandi/toshl-go"
	"github.com/stretchr/testify/assert"
)

type MockedHTTPClient struct {
	JSONString string
	Error      error
}

func (mc *MockedHTTPClient) Get(APIUrl string) (string, error) {
	return mc.JSONString, mc.Error
}

func TestClientDefaultURL(t *testing.T) {
	expected := "https://api.toshl.com"
	actual := toshl.DefaultBaseURL

	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestClientNewClient(t *testing.T) {
	c := toshl.NewClient("abcd1234", &MockedHTTPClient{})
	assert.NotNil(t, c.GetHTTPClient())
}

func TestClientAccounts(t *testing.T) {
	mc := &MockedHTTPClient{}
	mc.JSONString = `[{
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
	}]`

	c := toshl.NewClient("abcd1234", mc)
	accounts, _ := c.Accounts()
	assert.Len(t, accounts, 1)
}

func TestClientAccountsMultiple(t *testing.T) {
	mc := &MockedHTTPClient{}
	mc.JSONString = `[{
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
	},
	{
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
	}]`

	c := toshl.NewClient("abcd1234", mc)
	accounts, _ := c.Accounts()
	assert.Len(t, accounts, 2)
}

func TestClientGetAccount(t *testing.T) {
	mc := &MockedHTTPClient{}
	mc.JSONString = `{
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
	}`

	c := toshl.NewClient("abcd1234", mc)
	account, _ := c.GetAccount("42")
	assert.Equal(t, account.ID, "42")
}
