package controller

import (
	"net/http"

	"github.com/Devazt/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindPaslons(c *gin.Context) {

	var paslon []models.Paslon
	models.DB.Find(&paslon)
	c.JSON(http.StatusOK, gin.H{"message": "success", "paslon": paslon})
}

func FindPaslon(c *gin.Context) {

	var paslon models.Paslon
	id := c.Param("id")

	if err := models.DB.First(&paslon, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "paslon not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "paslon": paslon})
}

func CreatePaslon(c *gin.Context) {

	var paslon models.Paslon

	if err := c.ShouldBindJSON(&paslon); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&paslon)
	c.JSON(http.StatusOK, gin.H{"message": "create success", "paslon": paslon})

}

func UpdatePaslon(c *gin.Context) {

	var paslon models.Paslon
	id := c.Param("id")

	if err := c.ShouldBindJSON(&paslon); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&paslon).Where("id = ?", id).Updates(&paslon).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "paslon not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update success", "paslon": paslon})

}

func DeletePaslon(c *gin.Context) {

	var paslon models.Paslon
	id := c.Param("id")

	if err := models.DB.Delete(&paslon, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete success"})
}
