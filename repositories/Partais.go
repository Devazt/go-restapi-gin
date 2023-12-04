package repositories

import (
	"github.com/Devazt/go-restapi-gin/models"
	"gorm.io/gorm"
)

type PartaiRepo interface {
	FindPartais() ([]models.Partai, error)
	FindPartai(ID int) (models.Partai, error)
	CreatePartai(partai models.Partai) (models.Partai, error)
	UpdatePartai(partai models.Partai) (models.Partai, error)
	DeletePartai(partai models.Partai, ID int) (models.Partai, error)
}

type partaiRepo struct {
	db *gorm.DB
}

func RepoPartai(db *gorm.DB) *partaiRepo {
	return &partaiRepo{db}
}

func (r *partaiRepo) FindPartais() ([]models.Partai, error) {
	var partais []models.Partai
	err := r.db.Find(&partais).Error
	return partais, err
}

func (r *partaiRepo) FindPartai(ID int) (models.Partai, error) {
	var partai models.Partai
	err := r.db.First(&partai, ID).Error
	return partai, err
}

func (r *partaiRepo) CreatePartai(partai models.Partai) (models.Partai, error) {
	err := r.db.Create(&partai).Error
	return partai, err
}

func (r *partaiRepo) UpdatePartai(partai models.Partai) (models.Partai, error) {
	err := r.db.Save(&partai).Error
	return partai, err
}

func (r *partaiRepo) DeletePartai(partai models.Partai, ID int) (models.Partai, error) {
	err := r.db.Delete(&partai, ID).Error
	return partai, err
}
