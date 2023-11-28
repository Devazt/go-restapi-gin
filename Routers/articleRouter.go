package router

import (
	controller "github.com/Devazt/go-restapi-gin/controllers"
	"github.com/gin-gonic/gin"
)

func ArticleRoute(r *gin.RouterGroup) {

	r.GET("/article", controller.FindArticles)
	r.GET("/article/:id", controller.FindArticle)
	r.POST("/article", controller.CreateArticle)
	r.PUT("/article/:id", controller.UpdateArticle)
	r.DELETE("/article/:id", controller.DeleteArticle)

}
