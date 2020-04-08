package rest_errors

import (
	"net/http"
)

type RestErr struct {
	Message string  `json:"message"`
	Status 	int 	`json:"status"`
	Error 	string 	`json:"error"`
}

type logger interface {
	Error(string, error)
}

func NewBadRequestError(message string) *RestErr {
	error := RestErr{
		Message: message,
		Status: http.StatusBadRequest,
		Error: "bad_request",
	}

	return &error
}

func NewNotFoundError(message string) *RestErr {
	error := RestErr{
		Message: message,
		Status: http.StatusNotFound,
		Error: "not_found",
	}

	return &error
}

func NewUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status: http.StatusUnauthorized,
		Error: "unauthorized",
	}
}

func NewInternalServerError(message string, err error, logger logger) *RestErr {
	if logger != nil {
		logger.Error(message, err)
	}
	error := RestErr{
		Message: "Something went wrong.",
		Status: http.StatusInternalServerError,
		Error: "internal_server_error",
	}

	return &error
}