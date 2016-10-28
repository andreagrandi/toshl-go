package toshl

import (
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.toshl.com"
)

// Client handles API requests
type Client struct {
	client  *http.Client
	BaseURL *url.URL
}

// NewClient returns a new Toshl client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{client: httpClient, BaseURL: baseURL}

	return c
}
