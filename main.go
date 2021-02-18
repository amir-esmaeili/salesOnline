package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"personal/controllers"
	"personal/models"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}
	models.SetUpDataBase()
	routes := gin.Default()
	user := routes.Group("/user")
	{
		user.POST("/signup", controllers.SignUp)
		user.POST("/login", controllers.LogIn)
	}
	err := routes.Run(":8080")
	if err != nil {
		panic(err)
	}
}
