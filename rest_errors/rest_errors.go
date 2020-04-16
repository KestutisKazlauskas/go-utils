package rest_errors

import (
	"net/http"
	"encoding/json"
	"errors"
	"fmt"
)


type RestErr interface {
	Message() string
	Status() int
	Error() string
}

type restErr struct {
	ErrMessage string  `json:"message"`
	ErrStatus 	int 	`json:"status"`
	ErrError 	string 	`json:"error"`
}

func (err restErr) Message() string {
	return err.ErrMessage
}

func (err restErr) Status() int {
	return err.ErrStatus
}

func (err restErr) Error() string {
	return fmt.Sprintf(
		"message: %s, status: %d, error: %s", err.ErrMessage, err.ErrStatus, err.ErrError)
}

func NewRestError (message string, status int, error string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus: status,
		ErrError: error,
	}
}

func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr restErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

func NewBadRequestError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus: http.StatusBadRequest,
		ErrError: "bad_request",
	}
}

func NewNotFoundError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus: http.StatusNotFound,
		ErrError: "not_found",
	}
}

func NewUnauthorizedError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus: http.StatusUnauthorized,
		ErrError: "unauthorized",
	}
}

type logger interface {
	Error(string, error)
}

func NewInternalServerError(message string, err error, logger logger) RestErr {
	if logger != nil {
		logger.Error(message, err)
	}
	return restErr{
		ErrMessage: "Something went wrong.",
		ErrStatus: http.StatusInternalServerError,
		ErrError: "internal_server_error",
	}
}