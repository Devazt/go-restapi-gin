package models

import "time"

type Paslon struct {
	Id            int              `gorm:"primaryKey;autoIncrement" json:"id"`
	Name          string           `gorm:"type:varchar(100)" json:"name"`
	Serial        int              `gorm:"type:int;autoIncrement" json:"serial"`
	VisionMission string           `gorm:"type:text" json:"vision_mission"`
	Image         string           `gorm:"type:text" json:"image"`
	Partais       []PartaiResponse `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"partais"`
	Votes         []Vote           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"votes"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type PaslonVoteResponse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Serial int    `json:"serial"`
}

func (PaslonVoteResponse) TableName() string {
	return "paslons"
}
