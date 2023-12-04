package repositories

import (
	"github.com/Devazt/go-restapi-gin/models"
	"gorm.io/gorm"
)

type ArticleRepo interface {
	FindArticles() ([]models.Article, error)
	FindArticle(ID int) (models.Article, error)
	CreateArticle(article models.Article) (models.Article, error)
	UpdateArticle(article models.Article) (models.Article, error)
	DeleteArticle(article models.Article, ID int) (models.Article, error)
}

type articleRepo struct {
	db *gorm.DB
}

func RepoArticle(db *gorm.DB) *articleRepo {
	return &articleRepo{db}
}

func (r *articleRepo) FindArticles() ([]models.Article, error) {
	var articles []models.Article
	err := r.db.Find(&articles).Error
	return articles, err
}

func (r *articleRepo) FindArticle(ID int) (models.Article, error) {
	var article models.Article
	err := r.db.First(&article, ID).Error
	return article, err
}

func (r *articleRepo) CreateArticle(article models.Article) (models.Article, error) {
	err := r.db.Create(&article).Error
	return article, err
}

func (r *articleRepo) UpdateArticle(article models.Article) (models.Article, error) {
	err := r.db.Save(&article).Error
	return article, err
}

func (r *articleRepo) DeleteArticle(article models.Article, ID int) (models.Article, error) {
	err := r.db.Delete(&article, ID).Error
	return article, err
}
