package main

import (
	"github.com/gin-gonic/gin"
	"personal/controllers"
	"personal/models"
)

func main() {
	models.SetUpDataBase()
	routes := gin.Default()
	user := routes.Group("/user")
	{
		user.POST("/signup", controllers.SignUp)
	}
	err := routes.Run(":8080")
	if err != nil {
		panic(err)
	}
}
