package models

import "time"

type Article struct {
	ID        int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string `gorm:"type:varchar(255)" json:"title"`
	Image     string `gorm:"type:text" json:"image"`
	Content   string `gorm:"type:text" json:"content"`
	UserID    int    `json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ArticleResponse struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Image   string `json:"image"`
	Content string `json:"content"`
	UserID  int    `json:"user_id"`
}

func (ArticleResponse) TableName() string {
	return "articles"
}
