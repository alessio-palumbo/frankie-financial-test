package handlers

import (
	"github.com/alessio-palumbo/frankie-financial-test/cache"
	"github.com/gin-gonic/gin"
)

// Handler is a handler with an internal cache
type Handler struct {
	SessionCache *cache.SessionCache
}

// SetupRouter registers the listed endpoints and the returns a router.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Initialise the handler's cache to use to validate checkSessionKeys in this session.
	handler := Handler{cache.New()}

	r.POST("/isgood", handler.DeviceCheck)

	return r
}
