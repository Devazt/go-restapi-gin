package repositories

import (
	"github.com/Devazt/go-restapi-gin/models"
	"gorm.io/gorm"
)

type PaslonRepo interface {
	FindPaslons() ([]models.Paslon, error)
	FindPaslon(ID int) (models.Paslon, error)
	CreatePaslon(partai models.Paslon) (models.Paslon, error)
	UpdatePaslon(partai models.Paslon) (models.Paslon, error)
	DeletePaslon(partai models.Paslon, ID int) (models.Paslon, error)
}

type paslonRepo struct {
	db *gorm.DB
}

func RepoPaslon(db *gorm.DB) *paslonRepo {
	return &paslonRepo{db}
}

func (r *paslonRepo) FindPaslons() ([]models.Paslon, error) {
	var paslons []models.Paslon
	err := r.db.Find(&paslons).Error
	return paslons, err
}

func (r *paslonRepo) FindPaslon(ID int) (models.Paslon, error) {
	var paslon models.Paslon
	err := r.db.First(&paslon, ID).Error
	return paslon, err
}

func (r *paslonRepo) CreatePaslon(paslon models.Paslon) (models.Paslon, error) {
	err := r.db.Create(&paslon).Error
	return paslon, err
}

func (r *paslonRepo) UpdatePaslon(paslon models.Paslon) (models.Paslon, error) {
	err := r.db.Save(&paslon).Error
	return paslon, err
}

func (r *paslonRepo) DeletePaslon(paslon models.Paslon, ID int) (models.Paslon, error) {
	err := r.db.Delete(&paslon, ID).Error
	return paslon, err
}
