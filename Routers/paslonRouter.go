package router

import (
	controller "github.com/Devazt/go-restapi-gin/controllers"
	"github.com/gin-gonic/gin"
)

func PaslonRoute(r *gin.RouterGroup) {

	r.GET("/paslon", controller.FindPaslons)
	r.GET("/paslon/:id", controller.FindPaslon)
	r.POST("/paslon", controller.CreatePaslon)
	r.PUT("/paslon/:id", controller.UpdatePaslon)
	r.DELETE("/paslon/:id", controller.DeletePaslon)

}
