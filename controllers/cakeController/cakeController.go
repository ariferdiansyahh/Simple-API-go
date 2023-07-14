package cakecontroller

import (
	"errors"
	"net/http"

	"github.com/ariferdiansyah/ralali_backend/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "Bagian ini diperlukan"
	}
	return "Kesalahan yang tidak diketahui"
}

type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func Index(c *gin.Context) {
	var cakes []models.Cake
	models.DB.Order("rating desc, title").Find(&cakes)

	c.JSON(http.StatusOK, gin.H{"cakes": cakes, "message": "Berhasil"})
}
func Show(c *gin.Context) {
	var cake models.Cake
	id := c.Param("id")

	if err := models.DB.First(&cake, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"cake": cake})
}

func Create(c *gin.Context) {

	var cake models.Cake
	if err := c.ShouldBindJSON(&cake); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	models.DB.Create(&cake)
	c.JSON(http.StatusCreated, gin.H{"cakes": cake, "message": "Data Berhasil Ditambahkan"})
}

func Update(c *gin.Context) {
	var cake models.Cake
	id := c.Param("id")
	if err := c.ShouldBindJSON(&cake); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	if models.DB.Model(&cake).Where("id = ? ", id).Updates(&cake).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Gagal Update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil diperbarui"})
}

func Delete(c *gin.Context) {
	var cake models.Cake

	id := c.Param("id")
	if models.DB.Delete(&cake, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Gagal Delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data Berhasil Dihapus"})
}
