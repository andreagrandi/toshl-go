package toshl

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// DefaultBaseURL is ...
const (
	DefaultBaseURL = "https://api.toshl.com"
)

// Client handles API requests
type Client struct {
	client HTTPClient
}

// NewClient returns a new Toshl client
func NewClient(token string, httpClient HTTPClient) *Client {
	baseURL, _ := url.Parse(DefaultBaseURL)

	if httpClient == nil {
		httpClient = &RestHTTPClient{
			Client:  &http.Client{},
			BaseURL: baseURL.String(),
			Token:   token,
		}
	}

	c := &Client{client: httpClient}
	return c
}

// GetHTTPClient returns internal HTTPClient
func (c *Client) GetHTTPClient() HTTPClient {
	return c.client
}

// Accounts returns the list of Accounts
func (c *Client) Accounts(params *AccountQueryParams) ([]Account, error) {
	queryString := ""

	if params != nil {
		queryString = params.getQueryString()
	}

	res, err := c.client.Get("accounts", queryString)

	if err != nil {
		log.Fatal("GET /accounts/: ", err)
		return nil, err
	}

	var accounts []Account

	err = json.Unmarshal([]byte(res), &accounts)

	if err != nil {
		log.Fatalln("JSON: ", res)
		return nil, err
	}

	return accounts, nil
}

// GetAccount returns the a specific Account
func (c *Client) GetAccount(accountID string) (*Account, error) {
	res, err := c.client.Get(fmt.Sprintf("accounts/%s", accountID), "")

	if err != nil {
		log.Fatal(fmt.Sprintf("GET /accounts/%s: ", accountID), err)
		return nil, err
	}

	var account *Account

	err = json.Unmarshal([]byte(res), &account)

	if err != nil {
		log.Fatalln("JSON: ", res)
		return nil, err
	}

	return account, nil
}

// CreateAccount creates a Toshl Account
func (c *Client) CreateAccount(account *Account) error {
	jsonBytes, err := json.Marshal(account)

	if err != nil {
		log.Fatalln("CeateAccount: ", err)
		return err
	}

	jsonStr := string(jsonBytes)

	id, err := c.client.Post("accounts", jsonStr)

	if err != nil {
		log.Fatal("POST /accounts/ ", err)
		return err
	}

	account.ID = id

	return nil
}

// SearchAccount search for Account name and return an Account
func (c *Client) SearchAccount(accountName string) (*Account, error) {
	accounts, err := c.Accounts(nil)

	if err != nil {
		log.Fatal("GET /accounts/: ", err)
		return nil, err
	}

	for _, account := range accounts {
		if account.Name == accountName {
			return &account, nil
		}
	}

	return nil, nil
}

// UpdateAccount updates a Toshl Account
func (c *Client) UpdateAccount(account *Account) error {
	jsonBytes, err := json.Marshal(account)

	if err != nil {
		log.Fatalln("CeateAccount: ", err)
		return err
	}

	jsonStr := string(jsonBytes)

	accountResponse, err := c.client.Update(
		fmt.Sprintf("accounts/%s", account.ID), jsonStr)

	if err != nil {
		log.Fatal("PUT /accounts/ ", err)
		return err
	}

	err = json.Unmarshal([]byte(accountResponse), account)

	if err != nil {
		log.Fatalln("Cannot decode Account JSON")
		return err
	}

	return nil
}

// DeleteAccount deletes a Toshl Account
func (c *Client) DeleteAccount(account *Account) error {
	err := c.client.Delete(fmt.Sprintf("accounts/%s", account.ID))

	if err != nil {
		log.Fatal("DELETE /accounts/ ", err)
		return err
	}

	return nil
}
