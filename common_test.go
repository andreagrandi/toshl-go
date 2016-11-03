package toshl_test

import (
	"encoding/json"
	"testing"

	toshl "github.com/andreagrandi/toshl-go"
	"github.com/stretchr/testify/assert"
)

func TestCommonCurrencyDecode(t *testing.T) {
	var currency toshl.Currency
	currencyJSON := []byte(`{
        "code": "USD",
        "rate": 1,
        "fixed": false
    }`)

	err := json.Unmarshal(currencyJSON, &currency)

	if err != nil {
		t.Errorf("Cannot decode Currency JSON")
	}

	assert.Nil(t, err)
}

func TestCommonCurrencyDecodeError(t *testing.T) {
	var currency toshl.Currency
	currencyJSON := []byte(`{
        "code": 1,
        "rate": "fsdfa",
        "fixed": false
    }`)

	err := json.Unmarshal(currencyJSON, &currency)
	assert.NotNil(t, err)
}

func TestCommonMedianDecode(t *testing.T) {
	var median toshl.Median
	medianJSON := []byte(`{
        "expenses": 55,
        "incomes": 1300
    }`)

	err := json.Unmarshal(medianJSON, &median)

	if err != nil {
		t.Errorf("Cannot decode Median JSON")
	}

	assert.Nil(t, err)
}

func TestCommonMedianDecodeError(t *testing.T) {
	var median toshl.Median
	medianJSON := []byte(`{
        "expenses": "abc",
        "incomes": 1300
    }`)

	err := json.Unmarshal(medianJSON, &median)
	assert.NotNil(t, err)
}

func TestCommonGoalDecode(t *testing.T) {
	var goal toshl.Goal
	goalJSON := []byte(`{
        "amount": 63570,
        "start": "2013-07-01",
        "end": "2015-07-01"
    }`)

	err := json.Unmarshal(goalJSON, &goal)

	if err != nil {
		t.Errorf("Cannot decode Goal JSON")
	}

	assert.Nil(t, err)
}

func TestCommonGoalDecodeError(t *testing.T) {
	var goal toshl.Goal
	goalJSON := []byte(`{
        "amount": "abcd",
        "start": "2013-07-01",
        "end": "2015-07-01"
    }`)

	err := json.Unmarshal(goalJSON, &goal)
	assert.NotNil(t, err)
}
