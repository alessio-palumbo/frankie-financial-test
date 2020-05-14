package handlers

import (
	"errors"
	"net/http"
	"strings"

	e "github.com/alessio-palumbo/frankie-financial-test/errors"
	"github.com/alessio-palumbo/frankie-financial-test/models"
	"github.com/gin-gonic/gin"
)

// DeviceCheck validates the payload and returns a puppy if the request is correct.
// For invalid payloads it returns an error with a description of the issue found
func DeviceCheck(c *gin.Context) {
	var body []models.DeviceCheckDetails

	err := c.BindJSON(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, e.BadRequestJSONError())
		return
	}

	var errorsList []string
	for _, p := range body {
		err := validatePayload(p)
		if err != nil {
			errorsList = append(errorsList, err.Error())
		}
	}

	if len(errorsList) > 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, e.StringJSONError(strings.Join(errorsList, ", ")))
		return
	}

	c.JSON(200, gin.H{
		"puppy": true,
	})
}

func validatePayload(payload models.DeviceCheckDetails) error {

	if !isValidCheckType(payload.CheckType) {
		return errors.New("invalid checkType")
	}

	if !isValidActivityType(payload.ActivityType) {
		return errors.New("invalid activityType")
	}

	return nil
}
