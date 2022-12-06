package main

import (
	// gogin
	"flag"

	"github.com/gin-gonic/gin"
)

type data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	port := flag.String("port", "8080", "port to run the server on")
	flag.Parse()

	g := gin.Default()

	g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"service": "C",
			"message": "ServiceC",
			"data":    data{Name: "sam", Age: 30},
		})
	})

	g.Run(":" + *port)
}
