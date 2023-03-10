package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Ball struct {
	Name string
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	balls := []Ball{
		Ball{"bill"},
		Ball{"ted"},
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"balls":   balls,
		})
	})

	r.Run()
}
