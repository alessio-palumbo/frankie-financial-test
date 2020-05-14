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
	sampleRequest1 = `[{"checkType":"DEVICE","activityType":"","checkSessionKey":"",
		"activityData":[{"kvpKey":"","kvpValue":"","kvpType":""}]}]`
	sampleRequest2 = `[{"checkType":"TOOL","activityType":"","checkSessionKey":"",
		"activityData":[{"kvpKey":"","kvpValue":"","kvpType":""}]}]`

	invalidCheckTypeResponse = `{"code":500,"message":"invalid checkType"}`
)

func TestDeviceCheck(t *testing.T) {

	tests := []struct {
		name       string
		payload    string
		statusCode int
		response   string
	}{
		{
			name:       "Valid checkType",
			payload:    sampleRequest1,
			statusCode: 200,
			response:   successResponse,
		},
		{
			name:       "Empty request",
			payload:    "",
			statusCode: 400,
			response:   badRequestResponse,
		},
		{
			name:       "Invalid checkType",
			payload:    sampleRequest2,
			statusCode: 500,
			response:   invalidCheckTypeResponse,
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
