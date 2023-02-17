package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 后续希望可以同时支持离线建库以及在线检索
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
