package handlers

import (
	"net/http"

	"github.com/alessio-palumbo/frankie-financial-test/errors"
	"github.com/alessio-palumbo/frankie-financial-test/models"
	"github.com/gin-gonic/gin"
)

// DeviceCheck validates the payload and returns a puppy if the request is correct.
// For invalid payloads it returns an error with a description of the issue found
func DeviceCheck(c *gin.Context) {
	var body []models.DeviceCheckDetails

	err := c.BindJSON(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.BadRequestJSONError())
		return
	}

	c.JSON(200, gin.H{
		"puppy": true,
	})
}
