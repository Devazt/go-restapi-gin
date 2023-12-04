package main

import (
	"fmt"

	"github.com/Devazt/go-restapi-gin/database"
	"github.com/Devazt/go-restapi-gin/pkg/postgres"
	"github.com/Devazt/go-restapi-gin/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	postgres.DatabaseInit()
	database.RunMigration()

	router.RouteInit(r.Group("/api/v1"))

	fmt.Println("server running localhost:5000")
	r.Run(":5000")
}
