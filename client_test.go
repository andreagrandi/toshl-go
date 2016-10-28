package toshl

import (
	"testing"
)

func TestClientDefaultURL(t *testing.T) {
	expected := "https://api.toshl.com"
	actual := defaultBaseURL

	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}

func TestClientNewClient(t *testing.T) {
	expected := "https://api.toshl.com"
	c := NewClient(nil)
	actual := c.BaseURL.String()

	if actual != expected {
		t.Errorf("Test failed, expected: '%s', got:  '%s'", expected, actual)
	}
}
