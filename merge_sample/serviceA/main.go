package main

import (
	// gogin
	"flag"

	"github.com/gin-gonic/gin"
)

func main() {

	port := flag.String("port", "8080", "port to run the server on")
	flag.Parse()

	g := gin.Default()

	g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"service": "A",
			"message": "ServiceA",
		})
	})

	g.Run(":" + *port)
}
