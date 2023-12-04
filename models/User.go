package models

import "time"

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`
	Address   string    `gorm:"type:varchar(255)" json:"address"`
	Gender    string    `gorm:"type:varchar(255)" json:"gender"`
	Username  string    `gorm:"type:varchar(255)" json:"username"`
	Password  string    `gorm:"type:varchar(255)" json:"password"`
	Articles  []Article `json:"articles"`
	Vote      Vote      `json:"vote"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserVoteResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Gender  string `json:"gender"`
}

func (UserVoteResponse) TableName() string {
	return "users"
}
