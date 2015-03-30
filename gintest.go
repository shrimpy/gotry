package main

import (
    "os"
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "hello world")
    })
    router.GET("/ping", func(c *gin.Context) {
        c.String(http.StatusOK, "pong")
    })
    router.POST("/submit", func(c *gin.Context) {
        c.String(http.StatusUnauthorized, "not authorized")
    })
    router.PUT("/error", func(c *gin.Context) {
        c.String(http.StatusInternalServerError, "and error happened :(")
    })
    router.Run(":" + os.Getenv("HTTP_PLATFORM_PORT"))
}