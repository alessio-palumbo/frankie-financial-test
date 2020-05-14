package handlers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/alessio-palumbo/frankie-financial-test/cache"
	"github.com/alessio-palumbo/frankie-financial-test/enums"
	"github.com/alessio-palumbo/frankie-financial-test/models"
)

// isValidCheckType matches the given checkType against the list of accepted checkTypes.
func isValidCheckType(ct string) bool {
	if _, ok := enums.CheckTypes[ct]; ok {
		return true
	}
	return false
}

// isValidActivityType matches the given activityType against the list of accepted activityTypes.
// If the activityType starts with an underscore `_`, we assume it is a vendor specific
// activityType and accept the value.
func isValidActivityType(at string) bool {
	if _, ok := enums.ActivityTypes[at]; ok {
		return true
	}

	if strings.HasPrefix(at, "_") {
		return true
	}

	return false
}

// isValidCheckSessionKey makes sure that checkSessionKeys are unique within a session.
func isValidCheckSessionKey(sk string, cache *cache.SessionCache) bool {
	if cache.Has(sk) {
		return false
	}

	cache.Store(sk)
	return true
}

// isValidActivityData validates the following:
// * that a kpvKey hasn't already been used in current call
// * that the given kvpType matches the type of data described by kvpValue
// It returns an error comprising a list of any violation found or nil, if valid.
func isValidActivityData(kvPairs []models.KeyValuePair, usedKeys map[string]bool) error {
	var errorsList []string

	for _, p := range kvPairs {
		if _, found := usedKeys[p.KvpKey]; found {
			errorsList = append(errorsList, "duplicate kvpKey: "+p.KvpKey)
		}
		usedKeys[p.KvpKey] = true

		var err error

		switch p.KvpType {
		case enums.KVPGeneralString:
		case enums.KVPGeneralBool:
			_, err = strconv.ParseBool(p.KvpValue)
		case enums.KVPGeneralInteger:
			_, err = strconv.ParseInt(p.KvpValue, 10, 64)
		case enums.KVPGeneralFloat:
			_, err = strconv.ParseFloat(p.KvpValue, 64)
		default:
			errorsList = append(errorsList, fmt.Sprintf("unknown kvpType '%s' for '%s'", p.KvpType, p.KvpKey))
		}
		if err != nil {
			errorsList = append(errorsList, fmt.Sprintf("invalid kvpType '%s' for '%s'", p.KvpType, p.KvpKey))
		}

	}

	if len(errorsList) > 0 {
		return errors.New(strings.Join(errorsList, ", "))
	}

	return nil
}
