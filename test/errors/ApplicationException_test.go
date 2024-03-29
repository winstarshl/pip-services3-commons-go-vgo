package test_errors

import (
	"errors"
	"testing"

	cerrors "github.com/winstarshl/pip-services3-commons-go-vgo/errors"
	"github.com/stretchr/testify/assert"
)

func TestDefaultError(t *testing.T) {
	err := cerrors.NewError("")

	assert.Equal(t, "Unknown error", err.Message)
	assert.Equal(t, "UNKNOWN", err.Code)
	assert.Equal(t, 500, err.Status)
}

func TestWithCause(t *testing.T) {
	cause := errors.New("Cause error")
	err := cerrors.NewError("").WithCause(cause)

	assert.Equal(t, cause.Error(), err.Cause)
}

func TestWithCorrelationId(t *testing.T) {
	err := cerrors.NewError("").WithCorrelationId("123")

	assert.Equal(t, "123", err.CorrelationId)
}

func TestWithStatus(t *testing.T) {
	err := cerrors.NewError("").WithStatus(300)

	assert.Equal(t, 300, err.Status)
}
