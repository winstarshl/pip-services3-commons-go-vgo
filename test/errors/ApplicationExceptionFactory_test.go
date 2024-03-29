package test_errors

import (
	"testing"

	cerrors "github.com/winstarshl/pip-services3-commons-go-vgo/errors"
	"github.com/stretchr/testify/assert"
)

func TestCreateFromUnknown(t *testing.T) {
	d := &cerrors.ErrorDescription{
		Category:      cerrors.Unknown,
		CorrelationId: "123",
		Code:          "CODE",
		Message:       "Error message",
		Status:        321,
		Cause:         "Error cause",
		StackTrace:    "",
	}

	err := cerrors.ApplicationErrorFactory.Create(d)

	assert.Equal(t, cerrors.Unknown, err.Category)
	assert.Equal(t, "123", err.CorrelationId)
	assert.Equal(t, "CODE", err.Code)
	assert.Equal(t, "Error message", err.Message)
	assert.Equal(t, 500, err.Status)
	assert.Equal(t, "Error cause", err.Cause)
}

func TestCreateFromNotFound(t *testing.T) {
	d := &cerrors.ErrorDescription{
		Category:      cerrors.NotFound,
		CorrelationId: "123",
		Code:          "CODE",
		Message:       "Error message",
		Status:        321,
		Cause:         "Error cause",
		StackTrace:    "",
	}

	err := cerrors.ApplicationErrorFactory.Create(d)

	assert.Equal(t, cerrors.NotFound, err.Category)
	assert.Equal(t, "123", err.CorrelationId)
	assert.Equal(t, "CODE", err.Code)
	assert.Equal(t, "Error message", err.Message)
	assert.Equal(t, 404, err.Status)
	assert.Equal(t, "Error cause", err.Cause)
}
