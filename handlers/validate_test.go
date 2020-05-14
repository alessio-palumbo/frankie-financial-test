package handlers

import (
	"testing"

	"github.com/alessio-palumbo/frankie-financial-test/cache"
)

func Test_isValidCheckType(t *testing.T) {

	tests := []struct {
		name string
		ct   string
		want bool
	}{
		{
			name: "Valid checkType: DEVICE",
			ct:   "DEVICE",
			want: true,
		},
		{
			name: "Valid checkType: BIOMETRIC",
			ct:   "BIOMETRIC",
			want: true,
		},
		{
			name: "Valid checkType: COMBO",
			ct:   "COMBO",
			want: true,
		},
		{
			name: "Invalid checkType: device",
			ct:   "device",
			want: false,
		},
		{
			name: "Invalid checkType: SIGNATURE",
			ct:   "SIGNATURE",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidCheckType(tt.ct); got != tt.want {
				t.Errorf("isValidCheckType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidActivityType(t *testing.T) {

	tests := []struct {
		name string
		at   string
		want bool
	}{
		{
			name: "Valid activityType: SIGNUP",
			at:   "SIGNUP",
			want: true,
		},
		{
			name: "Valid activityType: LOGIN",
			at:   "LOGIN",
			want: true,
		},
		{
			name: "Valid activityType: PAYMENT",
			at:   "PAYMENT",
			want: true,
		},
		{
			name: "Valid activityType: CONFIRMATION",
			at:   "CONFIRMATION",
			want: true,
		},
		{
			name: "Valid activityType: vendor specific",
			at:   "_SIGNUP_VENDOR",
			want: true,
		},
		{
			name: "Invalid activityType: signup",
			at:   "signup",
			want: false,
		},
		{
			name: "Invalid activityType: COMPLAIN",
			at:   "COMPLAIN",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidActivityType(tt.at); got != tt.want {
				t.Errorf("isValidActivityType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidCheckSessionKey(t *testing.T) {

	tests := []struct {
		name         string
		sk           string
		existingKeys []string
		want         bool
	}{
		{
			name: "Empty cache",
			sk:   "CHECK_SESSION_KEY_001",
			want: true,
		},
		{
			name:         "Already in cache",
			sk:           "CHECK_SESSION_KEY_335",
			existingKeys: []string{"234", "334", "335"},
			want:         false,
		},
		{
			name:         "Not in cache",
			sk:           "CHECK_SESSION_KEY_111",
			existingKeys: []string{"234", "334", "335", "002", "134"},
			want:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := cache.New()
			for _, key := range tt.existingKeys {
				cache.Store("CHECK_SESSION_KEY_" + key)
			}

			if got := isValidCheckSessionKey(tt.sk, cache); got != tt.want {
				t.Errorf("isValidCheckSessionKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
