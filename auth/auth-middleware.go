package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Authenticate() gin.HandlerFunc {
	return func(context *gin.Context) {
		user, err := VerifyToken(context.Request)
		if err != nil {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Token not provided or expired"})
			context.Abort()
			return
		}
		context.Set("user", *user)
		context.Next()
	}
}