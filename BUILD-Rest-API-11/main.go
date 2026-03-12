package main

import (
	"example.com/BUILD-Rest-API-11/db"
	"example.com/BUILD-Rest-API-11/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") // localhost:8080
}
