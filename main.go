package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func sendPong(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Pong",
	})
}

func main() {
	router := gin.Default()
	router.GET("/ping", sendPong)

	port := os.Getenv("PORT")
	if port == "" {
		port = "7200"
	}

	router.Run("0.0.0.0:" + port)
}
