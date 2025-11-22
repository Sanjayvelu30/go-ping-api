package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() // Use Default() for logging and recovery middleware

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong update"})
	})

	// Run on custom port with error handling
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
