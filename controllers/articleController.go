package controller

import (
	"net/http"

	"github.com/Devazt/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindArticles(c *gin.Context) {

	var articles []models.Article
	models.DB.Find(&articles)
	c.JSON(http.StatusOK, gin.H{"message": "success", "articles": articles})
}

func FindArticle(c *gin.Context) {

	var article models.Article
	id := c.Param("id")

	if err := models.DB.First(&article, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "article not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "article": article})
}

func CreateArticle(c *gin.Context) {

	var article models.Article

	if err := c.ShouldBindJSON(&article); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&article)
	c.JSON(http.StatusOK, gin.H{"message": "create success", "data": article})
}

func UpdateArticle(c *gin.Context) {

	var article models.Article
	id := c.Param("id")

	if err := c.ShouldBindJSON(&article); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&article).Where("id = ?", id).Updates(&article).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "article not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "article updated", "data": article})

}

func DeleteArticle(c *gin.Context) {

	var article models.Article
	id := c.Param("id")

	if err := models.DB.Delete(&article, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "article deleted"})

}
