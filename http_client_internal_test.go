package toshl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestHTTPClientParseID(t *testing.T) {
	c := &RestHTTPClient{}
	id, _ := c.parseIDFromLocationHeader("https://api.toshl.com/accounts/42")
	assert.Equal(t, id, "42")
}
