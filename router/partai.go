package router

import (
	"github.com/Devazt/go-restapi-gin/handlers"
	"github.com/Devazt/go-restapi-gin/pkg/middlewares"
	"github.com/Devazt/go-restapi-gin/pkg/postgres"
	"github.com/Devazt/go-restapi-gin/repositories"
	"github.com/gin-gonic/gin"
)

func PartaiRoutes(e *gin.RouterGroup) {
	r := repositories.RepoPartai(postgres.DB)
	h := handlers.PartaiHandler(r)

	e.GET("/partai", h.FindPartais)
	e.GET("/partai/:id", h.FindPartai)
	e.POST("/partai", middlewares.UploadFile(h.CreatePartai))
	e.PUT("/partai/:id", middlewares.UploadFile(h.UpdatePartai))
	e.DELETE("/partai/:id", h.DeletePartai)
}
