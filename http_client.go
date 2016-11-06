package toshl

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// HTTPClient is an interface to define the client to access API resources
type HTTPClient interface {
	Get(APIUrl, queryString string) (string, error)
	Post(APIUrl, JSONPayload string) (string, error)
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

func (c *RestHTTPClient) setJSONContentTypeHeader(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
}

func (c *RestHTTPClient) getIDFromLocationHeader(
	response *http.Response) (string, error) {

	locationHeader := response.Header.Get("Location")

	id, err := c.parseIDFromLocationHeader(locationHeader)
	if err != nil {
		log.Fatal("Location URL parsing: ", err)
		return "", err
	}

	return id, nil
}

func (c *RestHTTPClient) parseIDFromLocationHeader(
	locationURL string) (string, error) {

	guid, err := url.Parse(locationURL)
	if err != nil {
		log.Fatal("Location URL parsing: ", err)
		return "", err
	}

	values := strings.Split(guid.Path, "/")

	if len(values) > 1 {
		id := values[len(values)-1]
		return id, nil
	}

	return "", errors.New("Cannot parse resource ID")
}

// Get takes an API endpoint and return a JSON string
func (c *RestHTTPClient) Get(APIUrl, queryString string) (string, error) {
	url := c.BaseURL + "/" + APIUrl

	if queryString != "" {
		url = url + "?" + queryString
	}

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

// Post takes an API endpoint and a JSON payload and return string ID
func (c *RestHTTPClient) Post(APIUrl, JSONPayload string) (string, error) {
	url := c.BaseURL + "/" + APIUrl
	var jsonStr = []byte(JSONPayload)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return "", err
	}

	// Set authorization token
	c.setAuthenticationHeader(req)

	// Set JSON content type
	c.setJSONContentTypeHeader(req)

	resp, err := c.Client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return "", err
	}

	defer resp.Body.Close()

	// Parse Location header to get ID
	id, err := c.getIDFromLocationHeader(resp)
	if err != nil {
		log.Fatal("Do: ", err)
		return "", err
	}

	return id, nil
}
