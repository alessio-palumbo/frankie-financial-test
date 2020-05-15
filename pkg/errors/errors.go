package errors

import (
	"net/http"
)

const (
	badRequestDefaultMessage = "payload is empty or malformed"
)

// CustomError returns an error struct with a statusCode and a message.
type CustomError struct {
	Code int `json:"code"`
	// Description of what went wrong (if we can tell)
	Message string `json:"message" example:"Everything is wrong. Go fix it."`
}

// New initialised a CustomError.
func New(statusCode int, message string) *CustomError {
	return &CustomError{
		Code:    statusCode,
		Message: message,
	}
}

// BadRequestJSONError returns a default response for an empty or malformed payload.
func BadRequestJSONError() *CustomError {
	return New(http.StatusBadRequest, badRequestDefaultMessage)
}

// StringJSONError returns the error string in JSON format.
func StringJSONError(err string) map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusInternalServerError,
		"message": err,
	}
}
