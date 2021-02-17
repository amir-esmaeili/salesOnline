package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	routes := gin.Default()
	routes.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})
	err := routes.Run()
	if err != nil {
		panic(err)
	}
}
