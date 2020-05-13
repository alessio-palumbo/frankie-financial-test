package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

const (
	successResponse    = `{"puppy":true}`
	badRequestResponse = `{"code":400,"message":"payload is empty or malformed"}`
)

var (
	sampleRequest = `[{"checkType":"dasfs","activityType":"","checkSessionKey":"",
		"activityData":[{"kvpKey":"","kvpValue":"","kvpType":""}]}]`
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
			payload:    sampleRequest,
			statusCode: 200,
			response:   successResponse,
		},
		{
			name:       "Empty request",
			payload:    "",
			statusCode: 400,
			response:   badRequestResponse,
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
