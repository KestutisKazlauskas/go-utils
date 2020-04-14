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
	message string  `json:"message"`
	status 	int 	`json:"status"`
	error 	string 	`json:"error"`
}

func (err restErr) Message() string {
	return err.message
}

func (err restErr) Status() int {
	return err.status
}

func (err restErr) Error() string {
	return fmt.Sprintf(
		"message: %s, status: %d, error: %s", err.message, err.status, err.error)
}

func NewRestError (message string, status int, error string) RestErr {
	return restErr{
		message: message,
		status: status,
		error: error,
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
		message: message,
		status: http.StatusBadRequest,
		error: "bad_request",
	}
}

func NewNotFoundError(message string) RestErr {
	return restErr{
		message: message,
		status: http.StatusNotFound,
		error: "not_found",
	}
}

func NewUnauthorizedError(message string) RestErr {
	return restErr{
		message: message,
		status: http.StatusUnauthorized,
		error: "unauthorized",
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
		message: "Something went wrong.",
		status: http.StatusInternalServerError,
		error: "internal_server_error",
	}
}