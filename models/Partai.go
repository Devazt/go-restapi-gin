package models

import "time"

type Partai struct {
	Id            int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string `gorm:"type:varchar(255)" json:"name"`
	Leader        string `gorm:"type:varchar(255)" json:"leader"`
	Serial        int    `gorm:"type:int;autoIncrement" json:"serial"`
	VisionMission string `gorm:"type:varchar(255)" json:"vision_mission"`
	Address       string `gorm:"type:varchar(255)" json:"address"`
	Image         string `gorm:"type:varchar(255)" json:"image"`
	PaslonID      int    `json:"paslon_id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type PartaiResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	PaslonID int    `json:"paslon_id"`
}

func (PartaiResponse) TableName() string {
	return "partais"
}
