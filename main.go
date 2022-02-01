package main

import (
	"net/http"
	
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	
	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s, this is version v0.0.0.1", name)
	})
	
	
	
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}