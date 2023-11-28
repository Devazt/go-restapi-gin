package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=localhost user=postgres password=root dbname=go_restapi_gin port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db

	db.AutoMigrate(&User{}, &Article{}, &Paslon{}, &Partai{}, &Vote{})

	fmt.Println("Connection Opened to Database")
}
