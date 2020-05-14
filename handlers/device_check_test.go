package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

const (
	successResponse       = `{"puppy":true}`
	badRequestResponse    = `{"code":400,"message":"payload is empty or malformed"}`
	cachedCheckSessionKey = "11111"
)

var (
	sampleRequest1 = `[{"checkType":"DEVICE","activityType":"LOGIN","checkSessionKey":"10001",
		"activityData":[{"kvpKey":"ip.address","kvpValue":"1.23.45.123","kvpType":"general.string"}]}]`
	sampleRequest2 = `[{"checkType":"TOOL","activityType":"LOGIN","checkSessionKey":"10002",
		"activityData":[{"kvpKey":"ip.address","kvpValue":"1.23.45.123","kvpType":"general.string"}]}]`
	sampleRequest3 = `[{"checkType":"COMBO","activityType":"LOGGING","checkSessionKey":"10003",
		"activityData":[{"kvpKey":"ip.address","kvpValue":"1.23.45.123","kvpType":"general.string"}]}]`
	sampleRequest4 = `[{"checkType":"COMBO","activityType":"LOGIN","checkSessionKey":"10001",
		"activityData":[{"kvpKey":"ip.address","kvpValue":"1.23.45.123","kvpType":"general.string"}]}]`
	sampleRequest5 = `[{"checkType":"BIOMETRIC","activityType":"PAYMENT","checkSessionKey":"10004",
		"activityData":[{"kvpKey":"ip.address","kvpValue":"1.23.45.123","kvpType":"general.string"}]},
		{"checkType":"DEVICE","activityType":"CONFIRMATION","checkSessionKey":"10005",
		"activityData":[{"kvpKey":"ip.address","kvpValue":"1.23.45.123","kvpType":"general.string"}]}]`
	sampleRequest6 = `[{"checkType":"COMBO","activityType":"SIGNUP","checkSessionKey":"10006",
		"activityData":[{"kvpKey":"ip.address","kvpValue":"1.23.45.123","kvpType":"general.integer"}]}]`

	invalidCheckTypeResponse       = `{"code":500,"message":"invalid checkType"}`
	invalidActivityTypeResponse    = `{"code":500,"message":"invalid activityType"}`
	invalidKvpTypeResponse         = `{"code":500,"message":"invalid kvpType for ip.address"}`
	invalidKvpKeyResponse          = `{"code":500,"message":"duplicate kvpKey: ip.address"}`
	invalidCheckSessionKeyResponse = `{"code":500,"message":"invalid checkSessionKey"}`
)

func TestDeviceCheck(t *testing.T) {

	tests := []struct {
		name       string
		payload    string
		statusCode int
		response   string
	}{
		{
			name:       "Valid request",
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
		{
			name:       "Invalid activityType",
			payload:    sampleRequest3,
			statusCode: 500,
			response:   invalidActivityTypeResponse,
		},
		{
			name:       "Invalid checkSessionKey",
			payload:    sampleRequest4,
			statusCode: 500,
			response:   invalidCheckSessionKeyResponse,
		},
		{
			name:       "Invalid KvpKey",
			payload:    sampleRequest5,
			statusCode: 500,
			response:   invalidKvpKeyResponse,
		},
		{
			name:       "Invalid activityData kpvType",
			payload:    sampleRequest6,
			statusCode: 500,
			response:   invalidKvpTypeResponse,
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
