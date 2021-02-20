package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"personal/forms"
	"personal/models"
)

func NewProduct(context *gin.Context) {
	userPtr, _ := context.Get("user")
	user := userPtr.(models.User)

	var newProductForm forms.NewProduct
	if err := context.ShouldBind(&newProductForm); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Please fill the required fields",
		})
		return
	}
	product, valid := newProductForm.Validator(context)
	if valid {
		product.Seller = user
		models.DB.Create(&product)
		context.JSON(http.StatusCreated, product)
		return
	}
	context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error": "Please fill the required fields",
	})
	return
}
