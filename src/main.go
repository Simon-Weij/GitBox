package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Static("/_app", "./frontend/build/_app")

	r.NoRoute(func(c *gin.Context) {
		c.File("./frontend/build/index.html")
	})

	r.Run()
}
