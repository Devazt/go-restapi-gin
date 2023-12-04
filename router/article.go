package router

import (
	"github.com/Devazt/go-restapi-gin/handlers"
	"github.com/Devazt/go-restapi-gin/pkg/middlewares"
	"github.com/Devazt/go-restapi-gin/pkg/postgres"
	"github.com/Devazt/go-restapi-gin/repositories"
	"github.com/gin-gonic/gin"
)

func ArticleRoutes(e *gin.RouterGroup) {
	r := repositories.RepoArticle(postgres.DB)
	h := handlers.ArticleHandler(r)

	e.GET("/articles", h.FindArticles)
	e.GET("/article/:id", h.FindArticle)
	e.POST("/article", middlewares.Auth(middlewares.UploadFile(h.CreateArticle)))
	e.PUT("/article/:id", middlewares.Auth(middlewares.UploadFile(h.UpdateArticle)))
	e.DELETE("/article/:id", h.DeleteArticle)
}
