package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	coffeeState := "off"

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the Coffee API")
	})

	router.GET("/api", func(c *gin.Context) {
		c.String(http.StatusOK, coffeeState)
	})

	// Change the state to turn on / off the coffee maker
	router.POST("/api/:trigger", func(c *gin.Context) {
		trigger := c.Param("trigger")

		if trigger == "on" || trigger == "off" {
			coffeeState = trigger
			c.String(http.StatusOK, "Success")
		} else {
			c.String(404, "Page not found")
		}
	})

	router.NoRoute(func(c *gin.Context) {
		c.String(404, "Page not found")
	})
	router.Run(":8080")
}
