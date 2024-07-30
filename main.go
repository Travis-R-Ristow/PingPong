package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func sendTestRsp(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hi, yes im here.",
	})
}

func main() {
	router := gin.Default()
	router.GET("/ping", sendTestRsp)

	port := os.Getenv("PORT")
	if port == "" {
		port = "7200"
	}

	router.Run("0.0.0.0:" + port)
}
