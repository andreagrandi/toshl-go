package toshl

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// HTTPClient is an interface to define the client to access API resources
type HTTPClient interface {
	Get(APIUrl string) (string, error)
}

// RestHTTPClient is a real implementation of the HTTPClient
type RestHTTPClient struct {
	BaseURL string
	Token   string
	Client  *http.Client
}

func (c *RestHTTPClient) setAuthenticationHeader(req *http.Request) {
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
}

// Get takes an API endpoint and return a JSON string
func (c *RestHTTPClient) Get(APIUrl string) (string, error) {
	url := c.BaseURL + "/" + APIUrl

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return "", err
	}

	// Set authorization token
	c.setAuthenticationHeader(req)

	resp, err := c.Client.Do(req)
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
