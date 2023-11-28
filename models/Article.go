package models

import "time"

type Article struct {
	Id        int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string `gorm:"type:varchar(100)" json:"title"`
	Author    string `gorm:"type:varchar(100)" json:"author"`
	Image     string `gorm:"type:text" json:"image"`
	Content   string `gorm:"type:text" json:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
