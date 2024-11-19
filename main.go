package main

import (
	"github.com/gin-gonic/gin"
	"gocourse.com/restapi/db"
	"gocourse.com/restapi/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}
