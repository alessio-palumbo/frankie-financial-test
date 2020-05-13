package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

const (
	ResponseSuccess = `{"puppy":true}`
)

func TestDeviceCheck(t *testing.T) {

	tests := []struct {
		name       string
		payload    string
		statusCode int
		response   string
	}{
		{
			name:       "Success",
			payload:    "",
			statusCode: 200,
			response:   ResponseSuccess,
		},
	}

	router := SetupRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/isgood", strings.NewReader(tt.payload))
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.statusCode, w.Code)
			assert.Equal(t, tt.response, w.Body.String())
		})
	}
}
