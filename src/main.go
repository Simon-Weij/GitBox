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

	r.Static("/_app", "./dist/_app")

	r.NoRoute(func(c *gin.Context) {
		c.File("./dist/index.html")
	})

	r.Run()
}
