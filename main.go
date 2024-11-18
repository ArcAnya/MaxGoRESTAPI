package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gocourse.com/restapi/db"
	"gocourse.com/restapi/models"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func getEvents(context *gin.Context) { // handler function for GET /events - using the gin context
	events := models.GetAllEvents() // calling the GetAllEvents function from the models package
	context.JSON(                   // sending a response back: status code + data (can by anything but often struct or map)
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

	event.Save() // saving the event to the database

	context.JSON(http.StatusCreated, gin.H{"message": "Event created successfully!", "event": event})

}
