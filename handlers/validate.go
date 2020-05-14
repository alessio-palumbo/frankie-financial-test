package handlers

import (
	"strings"

	"github.com/alessio-palumbo/frankie-financial-test/cache"
	"github.com/alessio-palumbo/frankie-financial-test/models"
)

func isValidCheckType(ct string) bool {
	if _, ok := models.CheckTypes[ct]; ok {
		return true
	}
	return false
}

func isValidActivityType(at string) bool {
	if _, ok := models.ActivityTypes[at]; ok {
		return true
	}

	if strings.HasPrefix(at, "_") {
		return true
	}

	return false
}

func isValidCheckSessionKey(sk string, cache *cache.SessionCache) bool {
	if cache.Has(sk) {
		return false
	}

	cache.Store(sk)
	return true
}
