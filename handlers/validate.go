package handlers

import "github.com/alessio-palumbo/frankie-financial-test/models"

func isValidCheckType(ct string) bool {
	if _, ok := models.CheckTypes[ct]; ok {
		return true
	}
	return false
}
