package router

import (
	"github.com/Devazt/go-restapi-gin/handlers"
	"github.com/Devazt/go-restapi-gin/pkg/postgres"
	"github.com/Devazt/go-restapi-gin/repositories"
	"github.com/gin-gonic/gin"
)

func UserRoutes(e *gin.RouterGroup) {
	r := repositories.RepoUser(postgres.DB)
	h := handlers.UserHandler(r)

	e.GET("/users", h.FindUsers)
	e.GET("/user/:id", h.FindUser)
	e.POST("/user", h.CreateUser)
}
