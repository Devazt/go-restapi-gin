package router

import (
	controller "github.com/Devazt/go-restapi-gin/controllers"
	"github.com/gin-gonic/gin"
)

func PartaiRoute(r *gin.RouterGroup) {

	r.GET("/partai", controller.FindPartais)
	r.GET("/partai/:id", controller.FindPartai)
	r.POST("/partai", controller.CreatePartai)
	r.PUT("/partai/:id", controller.UpdatePartai)
	r.DELETE("/partai/:id", controller.DeletePartai)

}
