package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
	_ = engine.Run(":8099")
}
