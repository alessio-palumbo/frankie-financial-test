package models

import "github.com/alessio-palumbo/frankie-financial-test/pkg/enums"

// DeviceCheckDetails contains any/all details we want to pass on to the device/biometric
// checking service as part of an activity/transaction. A transaction isn't just a payment,
// but represent a number of different interaction types.
type DeviceCheckDetails struct {
	// Describes the type of check service we need to verify with. Choices are:
	// * DEVICE: Services that will be checking device characteristics
	// * BIOMETRIC: Services that will be checking biomentric characteristics
	// * COMBO: If you're using a service that combines both device and biometric information, use this.
	CheckType string `json:"checkType" enums:"DEVICE,BIOMETRIC,COMBO"`

	// The type of activity we're checking. Choices are:
	// * SIGNUP: Used when an entity is signing up to your service
	// * LOGIN: Used when an already registered entity is logging in to your service
	// * PAYMENT: Used when you wish to check that all is well for a payment
	// * CONFIRMATION: User has confirmed an action and you wish to double check they're still legitimate
	//
	// You can also supply vendor specific activityTypes if you know them. To do this, make the first character an underscore _.
	// So for example, to use BioCatch's LOGIN_3 type, you can send "_LOGIN_3" as a value. Note, if you do this, there is no error checking on the Frankie side, and thus if you supply an incorrect value, the call will fail.
	ActivityType string `json:"activityType" enums:"SIGNUP,LOGIN,PAYMENT,CONFIRMATION"`

	// The unique session based ID that will be checked against the service.
	// Service key must be unique or an error will be returned.
	CheckSessionKey string `json:"checkSessionKey"`

	// A collection of loosely typed Key-Value-Pairs, which contain arbitrary data to be passed on to the verification services.
	// The API will verify that:
	// * the list of "Keys" provided are unique to the call (no double-ups)
	// * that the Value provided matches the Type specified.
	//
	// Should the verification fail, the error message returned will include information for each KVP pair that fails
	ActivityData []KeyValuePair `json:"activityData"`
}

// KeyValuePair contains arbitrary data to be passed on to the verification services.
type KeyValuePair struct {
	KvpKey   string `json:"kvpKey" example:"ip.address"`
	KvpValue string `json:"kvpValue" example:"1.23.45.123"`
	// Used to describe the contents of the KVP data.
	KvpType enums.KVPType `json:"kvpType" enums:"general.string,general.integer,general.float,general.bool"`
}
