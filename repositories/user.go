package repositories

import (
	"github.com/Devazt/go-restapi-gin/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	FindUsers() ([]models.User, error)
	FindUser(ID int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
}

type userRepo struct {
	db *gorm.DB
}

func RepoUser(db *gorm.DB) *userRepo {
	return &userRepo{db}
}

func (r *userRepo) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Articles").Find(&users).Error
	return users, err
}

func (r *userRepo) FindUser(ID int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Articles").First(&user, ID).Error
	return user, err
}

func (r *userRepo) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}
