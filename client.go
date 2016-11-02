package toshl

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.toshl.com"
)

// Client handles API requests
type Client struct {
	client HTTPClient
}

// NewClient returns a new Toshl client
func NewClient(token string, httpClient HTTPClient) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)

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
