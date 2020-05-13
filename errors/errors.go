package errors

import (
	"net/http"
)

const (
	badRequestDefaultMessage = "payload is empty or malformed"
)

// BadRequestJSONError returns a default response for an empty or malformed payload
func BadRequestJSONError() map[string]interface{} {
	return map[string]interface{}{
		"code":    http.StatusBadRequest,
		"message": badRequestDefaultMessage,
	}
}
