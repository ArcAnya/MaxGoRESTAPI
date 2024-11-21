package routes

import (
	"github.com/gin-gonic/gin"
	"gocourse.com/restapi/middlewares"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")

	authenticated.Use(middlewares.Authenticate)

	authenticated.POST("/events", middlewares.Authenticate, createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	// can also use middleware on a per route basis:
	// server.POST("/events", middlewares.Authenticate, createEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)

	server.Run(":8080")
}
