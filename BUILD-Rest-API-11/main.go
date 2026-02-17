package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// test
func main() {
	server := gin.Default()
	server.GET("/events", getEvents)
	server.Run(":8080") // listens on 0.0.0.0:8080 by default
}

func getEvents(Context *gin.Context) {
	Context.JSON(http.StatusOK, gin.H{"message": "Hello"})
}
