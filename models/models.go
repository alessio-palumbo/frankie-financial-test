package models

// DeviceCheckDetails contains any/all details we want to pass on to the device/biometric
// checking service as part of an activity/transaction. A transaction isn't just a payment,
// but represent a number of different interaction types.
type DeviceCheckDetails struct {
	CheckType       string         `json:"checkType"`
	ActivityType    string         `json:"activityType"`
	CheckSessionKey string         `json:"checkSessionKey"`
	ActivityData    []KeyValuePair `json:"activityData"`
}

// KeyValuePair contains arbitrary data to be passed on to the verification services.
type KeyValuePair struct {
	KvpKey   string  `json:"kvpKey"`
	KvpValue string  `json:"kvpValue"`
	KvpType  KVPType `json:"kvpType"`
}
