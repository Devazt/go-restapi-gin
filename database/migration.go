package database

import (
	"fmt"

	"github.com/Devazt/go-restapi-gin/models"
	"github.com/Devazt/go-restapi-gin/pkg/postgres"
)

func RunMigration() {
	err := postgres.DB.AutoMigrate(&models.Paslon{}, &models.Partai{}, &models.User{}, &models.Vote{}, &models.Article{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Migration Success")
}
