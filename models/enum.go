package models

// CheckType describes the type of check service we need to verify with.
type CheckType string

const (
	// CheckDevice describes a service checking device characteristics.
	CheckDevice CheckType = "DEVICE"

	// CheckBiometric describes a service checking biometric characteristics.
	CheckBiometric CheckType = "BIOMETRIC"

	// CheckCombo describes a service that combines both device and biometric information.
	CheckCombo CheckType = "COMBO"
)

// ActivityType is the type of activity we are checking.
// Vendor specific ActivityType can also be supplied when prepended with an underscore(_).
// In this case there is no error checking, and thus if the supplied value is incorrect
// the call will fail.
type ActivityType string

const (
	// ActivitySignup is used when an entity is signing up to the service.
	ActivitySignup ActivityType = "SIGNUP"

	// ActivityLogin is used to log in to the service by an already registered entity.
	ActivityLogin ActivityType = "LOGIN"

	// ActivityPayment is used to check that all is well for a payment.
	ActivityPayment ActivityType = "PAYMENT"

	// ActivityConfirmation is used to double check that a user is still legitimate when
	// they confirm an action.
	ActivityConfirmation ActivityType = "CONFIRMATION"
)

// KVPType describes the content of the KVP data.
type KVPType string

// KVPTypes for general types
const (
	KVPGeneralString  KVPType = "general.string"
	KVPGeneralInteger KVPType = "general.integer"
	KVPGeneralFloat   KVPType = "general.float"
	KVPGeneralBool    KVPType = "general.bool"
)
