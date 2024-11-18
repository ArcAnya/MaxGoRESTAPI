package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gocourse.com/restapi/db"
	"gocourse.com/restapi/models"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // new route to get a single event by ID
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get event"})
		return
	}

	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) { // handler function for GET /events - using the gin context
	events, err := models.GetAllEvents() // calling the GetAllEvents function from the models package
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get events"})
		return
	}

	context.JSON( // sending a response back: status code + data (can by anything but often struct or map)
		http.StatusOK, // status code
		events,        // map with a message key, ex: gin.H{"message": "Hello"}
	)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event) // binding the request body to the event struct.
	// Should make sure the request body is JSON & of type Event
	// ShouldBindJSON is pretty forgiving if data is missing => can make required by adding tags to the struct fields: `json:"title" binding:"required"`

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 1

	err = event.Save() // saving the event to the database
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to save event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully!", "event": event})

}
