package main

import (
	"github.com/gin-gonic/gin"
	"os"
  "fmt"
)

func sendPong(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Pong",
	})
}

func main() {
	router := gin.Default()

 // this wont work but we try 
  router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
    c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // Handle preflight request
			return
		}

		c.Next()
	})

	router.GET("/ping", sendPong)

	router.POST("/pong", func(c *gin.Context) {
    var oofs interface{}

    c.ShouldBindJSON(&oofs);

    fmt.Println("Package: ", oofs)
		// would save oofs to db i guess
		sendPong(c)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "7200"
	}

	router.Run("0.0.0.0:" + port)
}
