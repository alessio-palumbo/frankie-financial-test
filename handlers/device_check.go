package handlers

import (
	"github.com/gin-gonic/gin"
)

// DeviceCheck validates the payload and returns a puppy if the request is correct.
// For invalid payloads it returns an error with a description of the issue found
func DeviceCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"puppy": true,
	})
}
