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

func (mc *MockedHTTPClient) Get(APIUrl, queryString string) (string, error) {
	return mc.JSONString, mc.Error
}

func (mc *MockedHTTPClient) Post(APIUrl, JSONPayload string) (string, error) {
	return mc.JSONString, mc.Error
}

func (mc *MockedHTTPClient) Update(APIUrl, JSONPayload string) (string, error) {
	return mc.JSONString, mc.Error
}

func (mc *MockedHTTPClient) Delete(APIUrl string) error {
	return mc.Error
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
	accounts, _ := c.Accounts(nil)
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
	accounts, _ := c.Accounts(nil)
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

func TestClientCreateAccount(t *testing.T) {
	mc := &MockedHTTPClient{}
	mc.JSONString = "42"

	account := &toshl.Account{
		Name: "Test",
		Currency: toshl.Currency{
			Code: "GBP",
		},
	}

	c := toshl.NewClient("abcd1234", mc)
	c.CreateAccount(account)
	assert.Equal(t, account.ID, "42")
}

func TestClientSearchAccount(t *testing.T) {
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
		"id": "38",
		"name": "Tesla model A",
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
	account, _ := c.SearchAccount("Tesla model A")
	assert.Equal(t, account.ID, "38")
}

func TestClientUpdateAccount(t *testing.T) {
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
		"status": "active",
		"order": 0,
		"modified": "2012-09-04T13:55:15Z"
	}`

	account := &toshl.Account{
		Name: "Test",
		Currency: toshl.Currency{
			Code: "GBP",
		},
	}

	c := toshl.NewClient("abcd1234", mc)
	err := c.UpdateAccount(account)
	assert.Equal(t, account.Name, "Tesla model S")
	assert.Nil(t, err)
}

func TestClientDeleteAccount(t *testing.T) {
	mc := &MockedHTTPClient{}

	account := &toshl.Account{
		ID:   "50",
		Name: "Test",
		Currency: toshl.Currency{
			Code: "GBP",
		},
	}

	c := toshl.NewClient("abcd1234", mc)
	err := c.DeleteAccount(account)
	assert.Nil(t, err)
}

func TestClientMoveAccount(t *testing.T) {
	mc := &MockedHTTPClient{}

	account := &toshl.Account{
		ID:   "50",
		Name: "Test",
		Currency: toshl.Currency{
			Code: "GBP",
		},
	}

	c := toshl.NewClient("abcd1234", mc)
	err := c.MoveAccount(account, 10)
	assert.Nil(t, err)
}

func TestClientReorderAccounts(t *testing.T) {
	mc := &MockedHTTPClient{}

	order := &toshl.AccountsOrder{
		Order: []int{4, 12, 46, 2},
	}

	c := toshl.NewClient("abcd1234", mc)
	err := c.ReorderAccounts(order)
	assert.Nil(t, err)
}
