package router

import (
	controller "github.com/Devazt/go-restapi-gin/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoute(r *gin.RouterGroup) {

	r.POST("/auth/register", controller.Register)
	r.POST("/auth/login", controller.Login)

}
