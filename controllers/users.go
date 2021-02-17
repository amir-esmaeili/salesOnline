package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal/forms"
	"personal/models"
)




func SignUp(context *gin.Context)  {
	var signUpForm forms.SignUpForm
	if err := context.ShouldBindJSON(&signUpForm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "There is problem with the fields"})
		context.Abort()
		return
	}
	if err := models.DB.Where("email = ?", signUpForm.Email).
		First(&models.User{}).Error; err == nil {
		context.JSON(http.StatusConflict, gin.H{"message": "A user with this email already exists"})
		context.Abort()
		return
	}
	if signUpForm.Password != signUpForm.ConfirmPassword {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Passwords do not match"})
		context.Abort()
		return
	}
	hashedPassword, _ := models.HashPassword(signUpForm.Password)
	newUser := models.User{Name: signUpForm.Name, Email: signUpForm.Email, HashedPassword: hashedPassword}
	models.DB.Create(&newUser)
	context.JSON(http.StatusCreated, gin.H{"message": "New user created"})
}
