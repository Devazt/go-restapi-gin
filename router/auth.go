package router

import (
	"github.com/Devazt/go-restapi-gin/handlers"
	"github.com/Devazt/go-restapi-gin/pkg/postgres"
	"github.com/Devazt/go-restapi-gin/repositories"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(e *gin.RouterGroup) {
	r := repositories.RepoAuth(postgres.DB)
	h := handlers.AuthHandler(r)

	e.POST("/register", h.Register)
	e.POST("/login", h.Login)

}
