package rest_errors

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"net/http"
	"errors"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message", errors.New("database error"), nil)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
}

func TestNewBadRequestError(t *testing.T) {
	//TODO: Test!
}

func TestNewNotFoundError(t *testing.T) {
	//TODO: Test!
}

func TestNewError(t *testing.T) {
	//TODO: Test!
}