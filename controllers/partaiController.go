package controller

import (
	"net/http"

	"github.com/Devazt/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindPartais(c *gin.Context) {

	var partai []models.Partai
	models.DB.Find(&partai)
	c.JSON(http.StatusOK, gin.H{"message": "success", "partai": partai})
}

func FindPartai(c *gin.Context) {

	var partai models.Partai
	id := c.Param("id")

	if err := models.DB.First(&partai, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "partai not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "partai": partai})
}

func CreatePartai(c *gin.Context) {

	var partai models.Partai

	if err := c.ShouldBindJSON(&partai); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&partai)
	c.JSON(http.StatusOK, gin.H{"message": "create success", "data": partai})

}

func UpdatePartai(c *gin.Context) {

	var partai models.Partai
	id := c.Param("id")

	if err := c.ShouldBindJSON(&partai); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&partai).Where("id = ?", id).Updates(&partai).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "partai not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "update success", "data": partai})

}

func DeletePartai(c *gin.Context) {

	var partai models.Partai
	id := c.Param("id")

	if err := models.DB.Delete(&partai, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "delete success", "data": partai})

}
