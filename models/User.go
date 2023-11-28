package models

import "time"

type User struct {
	Id        int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string `gorm:"type:varchar(100)" json:"name"`
	Address   string `gorm:"type:varchar(300)" json:"address"`
	Gender    string `gorm:"type:varchar(100)" json:"gender"`
	Username  string `gorm:"type:varchar(100)" json:"username"`
	Password  string `gorm:"type:varchar(100)" json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
