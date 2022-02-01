package main

import (
    "net/http"
    
    "github.com/gin-gonic/gin"
)


func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*.tmpl")
    // This handler will match /user/john but will not match /user/ or /user
    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title": "Main website",
        })
    })
    
    
    
    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}