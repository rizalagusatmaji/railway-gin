package main

import (
	"gin/handler"
	"gin/middleware"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Use(middleware.ValidateAPIKey())
	r.GET("/login", handler.Login)
	// r.Use(middleware.ValidateUserToken())
	// r.GET("/recipes", handler.HandlerGetRecipe)

	r.Run()
}
