package drs

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlotStatus(t *testing.T) {
	ConfigSetup()
	badInput := SlotStatus{
		ExpectedReplenishmentDate: "2003-12-28 19:00:00",
		LastUseDate:               "1985-01-13 05:05:04",
	}
	goodInput := SlotStatus{
		RemainingQuantityInUnit:   3.5,
		OriginalQuantityInUnit:    10.0,
		TotalQuantityOnHand:       0,
		ExpectedReplenishmentDate: "2015-12-28T10:00:00Z",
		LastUseDate:               "2015-12-21T10:00:00Z",
	}
	result, err := ReportSlotStatus("", "", &goodInput)
	assert.NotNil(t, err)
	assert.False(t, result)
	if apiError, ok := err.(*APIError); ok {
		assert.Equal(t, http.StatusBadRequest, apiError.Code)
	}

	result, err = ReportSlotStatus("TEST", "", &goodInput)
	assert.NotNil(t, err)
	assert.False(t, result)
	if apiError, ok := err.(*APIError); ok {
		assert.Equal(t, http.StatusBadRequest, apiError.Code)
	}

	result, err = ReportSlotStatus("TEST", "TEST", &badInput)
	assert.NotNil(t, err)
	assert.False(t, result)
	if apiError, ok := err.(*APIError); ok {
		assert.Equal(t, http.StatusBadRequest, apiError.Code)
	}

	//make expected fine but last still bad
	badInput.ExpectedReplenishmentDate = "2015-12-28T10:00:00Z"
	result, err = ReportSlotStatus("TEST", "TEST", &badInput)
	assert.NotNil(t, err)
	assert.False(t, result)
	if apiError, ok := err.(*APIError); ok {
		assert.Equal(t, http.StatusBadRequest, apiError.Code)
	}

	result, err = ReportSlotStatus("TEST", "TEST", &goodInput)
	assert.Nil(t, err)
	assert.True(t, result)

	result, err = ReportSlotStatus("random", "random", &goodInput)
	assert.NotNil(t, err)
	assert.False(t, result)
}
