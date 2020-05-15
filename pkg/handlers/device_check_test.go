package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestDeviceCheck(t *testing.T) {

	tests := []struct {
		name       string
		payload    string
		statusCode int
		response   string
	}{
		{
			name: "Valid request",
			payload: `[{"checkType":"DEVICE","activityType":"LOGIN","checkSessionKey":"10001","activityData":
				[{"kvpKey":"ip.address","kvpValue":"1.23.45.123","kvpType":"general.string"}]}]`,
			statusCode: 200,
			response:   `{"puppy":true}`,
		},
		{
			name:       "Empty request",
			payload:    "",
			statusCode: 400,
			response:   `{"code":400,"message":"payload is empty or malformed"}`,
		},
		{
			name: "Invalid checkType",
			payload: `[{"checkType":"TOOL","activityType":"LOGIN","checkSessionKey":"10002","activityData":
				[{"kvpKey":"ip.address","kvpValue":"1.23.45.123","kvpType":"general.string"}]}]`,
			statusCode: 500,
			response:   `{"code":500,"message":"invalid checkType: TOOL"}`,
		},
		{
			name: "Invalid activityType",
			payload: `[{"checkType":"COMBO","activityType":"LOGGING","checkSessionKey":"10003","activityData":
				[{"kvpKey":"ip.address","kvpValue":"1.23.45.123","kvpType":"general.string"}]}]`,
			statusCode: 500,
			response:   `{"code":500,"message":"invalid activityType: LOGGING"}`,
		},
		{
			name: "Invalid checkSessionKey",
			payload: `[{"checkType":"COMBO","activityType":"LOGIN","checkSessionKey":"10001","activityData":
				[{"kvpKey":"ip.address","kvpValue":"1.23.45.123","kvpType":"general.string"}]}]`,
			statusCode: 500,
			response:   `{"code":500,"message":"invalid checkSessionKey: 10001"}`,
		},
		{
			name: "Invalid KvpKey",
			payload: `[{"checkType":"BIOMETRIC","activityType":"PAYMENT","checkSessionKey":"10004","activityData":
				[{"kvpKey":"ip.address","kvpValue":"1.23.45.123","kvpType":"general.string"}]},
				{"checkType":"DEVICE","activityType":"CONFIRMATION","checkSessionKey":"10005","activityData":
				[{"kvpKey":"ip.address","kvpValue":"1.23.45.123","kvpType":"general.string"}]}]`,
			statusCode: 500,
			response:   `{"code":500,"message":"duplicate kvpKey: ip.address"}`,
		},
		{
			name: "Invalid activityData kpvType",
			payload: `[{"checkType":"COMBO","activityType":"SIGNUP","checkSessionKey":"10006","activityData":
				[{"kvpKey":"ip.address","kvpValue":"1.23.45.123","kvpType":"general.integer"}]}]`,
			statusCode: 500,
			response:   `{"code":500,"message":"invalid kvpType 'general.integer' for 'ip.address'"}`,
		},
		{
			name: "Multiple invalid fields",
			payload: `[{"checkType":"BILLING","activityType":"PRICING","checkSessionKey":"10001","activityData":
				[{"kvpKey":"ip.address","kvpValue":"1.23.45.123","kvpType":"general.integer"}]}]`,
			statusCode: 500,
			response: `{"code":500,"message":"invalid checkType: BILLING, invalid activityType: PRICING, ` +
				`invalid checkSessionKey: 10001, invalid kvpType 'general.integer' for 'ip.address'"}`,
		},
		{
			name: "Multiple objects invalid fields",
			payload: `[{"checkType":"COMBO","activityType":"LOGIN","checkSessionKey":"10007","activityData":
				[{"kvpKey":"ip.address","kvpValue":"1.23.45.123","kvpType":"general.string"}]},
				{"checkType":"SCAN","activityType":"CHECK","checkSessionKey":"10007","activityData":
				[{"kvpKey":"ip.address","kvpValue":"21.23.34.122","kvpType":"general.integer"}]}]`,
			statusCode: 500,
			response: `{"code":500,"message":"invalid checkType: SCAN, invalid activityType: CHECK, invalid ` +
				`checkSessionKey: 10007, duplicate kvpKey: ip.address, invalid kvpType 'general.integer' for 'ip.address'"}`,
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
