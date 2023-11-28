package models

import "time"

type Vote struct {
	Id        int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
