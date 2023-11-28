package main

import (
	router "github.com/Devazt/go-restapi-gin/Routers"
	"github.com/Devazt/go-restapi-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDB()

	router.RouteInit(r.Group("/api/v1"))

	r.Run("localhost:5000")
}
