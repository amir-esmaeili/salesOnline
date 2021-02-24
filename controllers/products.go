package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"personal/forms"
	"personal/models"
	"strconv"
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
		product.SellerID = user.ID
		models.DB.Create(&product)
		context.JSON(http.StatusCreated, product)
		return
	}
	context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error": "Please fill the required fields",
	})
	return
}

func GetSellerProducts(context *gin.Context)  {
	sellerId := context.Param("seller_id")
	pageParam := context.DefaultQuery("page", "0")
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Page number is not valid",
		})
	}
	var products []models.Product
	models.DB.Where("seller_id = ?", sellerId).Find(&products)
	if page != 0 {
		postPerPage, _ := strconv.Atoi(os.Getenv("POST_PER_PAGE"))
		context.JSON(http.StatusOK, products[postPerPage* (page - 1): postPerPage* page])
		return
	}
	context.JSON(http.StatusOK, products)
}