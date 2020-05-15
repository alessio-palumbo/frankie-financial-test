package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	e "github.com/alessio-palumbo/frankie-financial-test/pkg/errors"
	"github.com/alessio-palumbo/frankie-financial-test/pkg/models"
)

type Success struct {
	// Everyone gets a puppy if the SDK output is good.
	Puppy bool `json:"puppy" example:"true"`
}

// DeviceCheck validates the payload and returns a puppy if the request is correct.
// For invalid payloads it returns an error with a description of the issue found
//
// @Summary Simple check to see if the service is running smoothly.
// @Description Validate the payload and returns a puppy if successful.
// @Description Otherwise it returns a status 500 error with a list of the violations.
// @Accept json
// @Produce json
// @Param deviceCheckDetails body []models.DeviceCheckDetails true "An array of objects that contain the details from each different provider wrapped up in the Universal SDK."
// @Success 200 {object} handlers.Success "The data is fine. No issues, and everyone gets a puppy."
// @Failure 500 {object} errors.CustomError "The system is presently unavailable, or running in a severely degraded state. Check the error message for details"
// @Router /isgood [post]
func (h *Handler) DeviceCheck(c *gin.Context) {
	var body []models.DeviceCheckDetails

	err := c.BindJSON(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, e.BadRequestJSONError())
		return
	}

	var errorsList []string
	activityDataKeys := make(map[string]bool)
	for _, p := range body {
		h.validatePayload(p, activityDataKeys, &errorsList)
	}

	if len(errorsList) > 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, e.StringJSONError(strings.Join(errorsList, ", ")))
		return
	}

	c.JSON(200, gin.H{
		"puppy": true,
	})
}

func (h *Handler) validatePayload(payload models.DeviceCheckDetails, activityDataKeys map[string]bool, errorsList *[]string) {

	if !isValidCheckType(payload.CheckType) {
		*errorsList = append(*errorsList, "invalid checkType: "+payload.CheckType)
	}

	if !isValidActivityType(payload.ActivityType) {
		*errorsList = append(*errorsList, "invalid activityType: "+payload.ActivityType)
	}

	if !isValidCheckSessionKey(payload.CheckSessionKey, h.SessionCache) {
		*errorsList = append(*errorsList, "invalid checkSessionKey: "+payload.CheckSessionKey)
	}

	if err := isValidActivityData(payload.ActivityData, activityDataKeys); err != nil {
		*errorsList = append(*errorsList, err.Error())
	}
}
