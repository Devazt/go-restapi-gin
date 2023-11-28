package controller

import (
	"net/http"

	"github.com/Devazt/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})

}

func Login(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Where("username = ? AND password = ?", user.Username, user.Password).First(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})

}
