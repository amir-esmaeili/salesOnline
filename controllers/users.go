package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal/auth"
	"personal/forms"
	"personal/models"
)




func SignUp(context *gin.Context)  {
	var signUpForm forms.SignUpForm
	if err := context.ShouldBindJSON(&signUpForm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Please fill the required fields"})
		context.Abort()
		return
	}
	if err := models.DB.Where("email = ?", signUpForm.Email).
		First(&models.User{}).Error; err == nil {
		context.JSON(http.StatusConflict, gin.H{"error": "A user with this email already exists"})
		context.Abort()
		return
	}
	if signUpForm.Password != signUpForm.ConfirmPassword {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		context.Abort()
		return
	}
	hashedPassword, _ := models.HashPassword(signUpForm.Password)
	newUser := models.User{Name: signUpForm.Name, Email: signUpForm.Email, HashedPassword: hashedPassword}
	models.DB.Create(&newUser)
	context.JSON(http.StatusCreated, gin.H{"error": "New user created"})
}

func LogIn(context *gin.Context)  {
	var loginForm forms.LogInForm
	if err := context.ShouldBindJSON(&loginForm); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Please fill the required fields"})
		context.Abort()
		return
	}
	var user models.User
	if err := models.DB.Where("Email = ?", loginForm.Email).First(&user).Error; err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Email or password is incorrect"})
		context.Abort()
		return
	}
	// If not verified
	if verified := user.CheckPasswordHash(loginForm.Password); !verified {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Email or password is incorrect"})
		context.Abort()
		return
	}
	token, err := auth.CreateToken(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something bad happened"})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": token})
}
