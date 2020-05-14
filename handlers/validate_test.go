package handlers

import (
	"testing"

	"github.com/alessio-palumbo/frankie-financial-test/cache"
	"github.com/alessio-palumbo/frankie-financial-test/models"
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

func Test_isValidActivityData(t *testing.T) {

	tests := []struct {
		name     string
		kvPairs  []models.KeyValuePair
		usedKeys map[string]bool
		wantErr  bool
	}{
		{
			name: "Valid keys",
			kvPairs: []models.KeyValuePair{
				models.KeyValuePair{
					KvpKey:   "ip.address",
					KvpValue: "127.23.12.12",
					KvpType:  "general.string",
				},
				models.KeyValuePair{
					KvpKey:   "id.user",
					KvpValue: "20435",
					KvpType:  "general.integer",
				},
				models.KeyValuePair{
					KvpKey:   "user.connected",
					KvpValue: "true",
					KvpType:  "general.bool",
				},
				models.KeyValuePair{
					KvpKey:   "interest.rate",
					KvpValue: "0.15",
					KvpType:  "general.float",
				},
			},
			usedKeys: make(map[string]bool),
			wantErr:  false,
		},
		{
			name: "Duplicate keys in same object",
			kvPairs: []models.KeyValuePair{
				models.KeyValuePair{
					KvpKey:   "ip.address",
					KvpValue: "127.23.12.12",
					KvpType:  "general.string",
				},
				models.KeyValuePair{
					KvpKey:   "ip.address",
					KvpValue: "120.12.23.1",
					KvpType:  "general.string",
				},
			},
			usedKeys: make(map[string]bool),
			wantErr:  true,
		},
		{
			name: "Duplicate keys in different objects",
			kvPairs: []models.KeyValuePair{
				models.KeyValuePair{
					KvpKey:   "ip.address",
					KvpValue: "127.23.12.12",
					KvpType:  "string",
				},
			},
			usedKeys: map[string]bool{
				"ip.address": true,
			},
			wantErr: true,
		},
		{
			name: "Duplicate keys in different objects",
			kvPairs: []models.KeyValuePair{
				models.KeyValuePair{
					KvpKey:   "ip.address",
					KvpValue: "127.23.12.12",
					KvpType:  "string",
				},
			},
			usedKeys: map[string]bool{
				"ip.address": true,
			},
			wantErr: true,
		},
		{
			name: "Invalid kvpType integer",
			kvPairs: []models.KeyValuePair{
				models.KeyValuePair{
					KvpKey:   "ip.address",
					KvpValue: "127.23.12.12",
					KvpType:  "general.integer",
				},
			},
			usedKeys: make(map[string]bool),
			wantErr:  true,
		},
		{
			name: "Invalid kvpType float",
			kvPairs: []models.KeyValuePair{
				models.KeyValuePair{
					KvpKey:   "mac.address",
					KvpValue: "00:A0:C9:14:C8:29",
					KvpType:  "general.float",
				},
			},
			usedKeys: make(map[string]bool),
			wantErr:  true,
		},
		{
			name: "Invalid kvpType bool",
			kvPairs: []models.KeyValuePair{
				models.KeyValuePair{
					KvpKey:   "interest.rate",
					KvpValue: "0.25",
					KvpType:  "general.bool",
				},
			},
			usedKeys: make(map[string]bool),
			wantErr:  true,
		},
		{
			name: "Unknown kvpType",
			kvPairs: []models.KeyValuePair{
				models.KeyValuePair{
					KvpKey:   "week.day",
					KvpValue: "mon,tue,wed,thu,fri,sat,sun",
					KvpType:  "general.enum",
				},
			},
			usedKeys: make(map[string]bool),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := isValidActivityData(tt.kvPairs, tt.usedKeys); (err != nil) != tt.wantErr {
				t.Errorf("isValidActivityData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
