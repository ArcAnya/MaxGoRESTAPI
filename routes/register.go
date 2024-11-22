package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gocourse.com/restapi/models"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get event"})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register for event"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registered for event successfully!"})

}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	var event models.Event
	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to cancel registration"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration cancelled successfully!"})
}
