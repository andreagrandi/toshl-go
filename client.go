package toshl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	token   string
}

// NewClient returns a new Toshl client
func NewClient(token string) *Client {
	httpClient := &http.Client{}
	baseURL, _ := url.Parse(defaultBaseURL)

	c := &Client{client: httpClient, BaseURL: baseURL, token: token}

	return c
}

func (c *Client) setAuthenticationHeader(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
}

func (c *Client) get(API string) (string, error) {
	url := c.BaseURL.String() + "/" + API

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return "", err
	}

	// Set authorization token
	c.setAuthenticationHeader(req)

	resp, err := c.client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return "", err
	}

	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ReadAll: ", err)
		return "", err
	}

	return string(bs), nil
}

// Accounts returns the list of Accounts
func (c *Client) Accounts() ([]Account, error) {
	res, err := c.get("accounts")

	if err != nil {
		log.Fatal("GET /accounts/: ", err)
		return nil, err
	}

	var account []Account

	err = json.Unmarshal([]byte(res), &account)

	if err != nil {
		log.Fatalln("JSON: ", res)
		return nil, err
	}

	return account, nil
}
