package forms

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"mime/multipart"
	"os"
	"path/filepath"
	"personal/models"
)

type NewProduct struct {
	Name      string                `form:"name" binding:"required"`
	Image     *multipart.FileHeader `form:"image"`
	Price     uint                  `form:"price" binding:"required"`
	Available uint                  `form:"available" binding:"required"`
	Off       uint                  `form:"off" gorm:"default:0"`
}

func (p NewProduct) Validator(context *gin.Context) (*models.Product, bool) {
	product := models.Product{
		Price: p.Price,
		Off: p.Off,
		Available: p.Available,
	}
	if len(p.Name) < 3 {
		return nil, false
	}
	product.Name = p.Name
	if p.Image != nil {
		ext := filepath.Ext(p.Image.Filename)
		name := uuid.New().String() + ext
		path := os.Getenv("MEDIA_ROOT") + name
		if err := context.SaveUploadedFile(p.Image, path); err != nil {
			return nil, false
		}
		product.Image = path
	}
	return &product, true
}
