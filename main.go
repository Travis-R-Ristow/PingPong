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

	router.GET("/ping", sendPong)

	router.GET("/pong", func(c *gin.Context) {
		oofs := c.Query("oofs")
    fmt.Println("Package: ", oofs)
		// save oofs to db i guess
		sendPong(c)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "7200"
	}

	router.Run("0.0.0.0:" + port)
}
