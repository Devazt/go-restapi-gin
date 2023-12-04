package repositories

import (
	"github.com/Devazt/go-restapi-gin/models"
	"gorm.io/gorm"
)

type AuthRepo interface {
	Login(username string) (models.User, error)
	Register(user models.User) (models.User, error)
}

type authRepo struct {
	db *gorm.DB
}

func RepoAuth(db *gorm.DB) *authRepo {
	return &authRepo{db}
}

func (r *authRepo) Login(username string) (models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *authRepo) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}
