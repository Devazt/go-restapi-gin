package router

import (
	"github.com/Devazt/go-restapi-gin/handlers"
	"github.com/Devazt/go-restapi-gin/pkg/middlewares"
	"github.com/Devazt/go-restapi-gin/pkg/postgres"
	"github.com/Devazt/go-restapi-gin/repositories"
	"github.com/gin-gonic/gin"
)

func PaslonRoutes(e *gin.RouterGroup) {
	r := repositories.RepoPaslon(postgres.DB)
	h := handlers.PaslonHandler(r)

	e.GET("/paslon", h.FindPaslons)
	e.GET("/paslon/:id", h.FindPaslon)
	e.POST("/paslon", middlewares.UploadFile(h.CreatePaslon))
	e.PUT("/paslon/:id", middlewares.UploadFile(h.UpdatePaslon))
	e.DELETE("/paslon/:id", h.DeletePaslon)
}
