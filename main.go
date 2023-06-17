package main

import (
	"fmt"
	"gin/handler"
	"gin/middleware"
	"os"

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
		fmt.Println(os.Getenv("APIKEY"))
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Use(middleware.ValidateAPIKey())
	r.POST("/login", handler.Login)
	// r.Use(middleware.ValidateUserToken())
	// r.GET("/recipes", handler.HandlerGetRecipe)

	r.Run()
}
