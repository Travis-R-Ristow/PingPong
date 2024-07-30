package main

import (
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

	router.Run("localhost:7200")
}
