package toshl

import (
	"encoding/json"
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
func (c *Client) Accounts() ([]Account, error) {
	res, err := c.client.Get("accounts")

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
