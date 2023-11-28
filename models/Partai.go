package models

import "time"

type Partai struct {
	Id            int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string `gorm:"type:varchar(100)" json:"name"`
	Leader        string `gorm:"type:varchar(100)" json:"leader"`
	Serial        int64  `gorm:"type:int;autoIncrement" json:"serial"`
	VisionMission string `gorm:"type:text" json:"vision_mission"`
	Image         string `gorm:"type:text" json:"image"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
