package handlers

import "github.com/gin-gonic/gin"

// SetupRouter registers the listed endpoints and the return a router.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/isgood", DeviceCheck)

	return r
}
