package partaisdto

type CreatePartaiReq struct {
	Name          string `json:"name" form:"name" validate:"required"`
	Leader        string `json:"leader" form:"leader" validate:"required"`
	VisionMission string `json:"vision_mission" form:"vision_mission" validate:"required"`
	Address       string `json:"address" form:"address" validate:"required"`
	Image         string `json:"image" form:"image" validate:"required"`
	PaslonID      string `json:"paslon_id" form:"paslon_id" validate:"required"`
}

type UpdatePartaiReq struct {
	Name          string `json:"name" form:"name"`
	Leader        string `json:"leader" form:"leader"`
	VisionMission string `json:"vision_mission" form:"vision_mission"`
	Address       string `json:"address" form:"address"`
	Image         string `json:"image" form:"image"`
	PaslonID      string `json:"paslon_id" form:"paslon_id"`
}

type PartaiRes struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Leader        string `json:"leader"`
	Serial        int    `json:"serial"`
	VisionMission string `json:"vision_mission"`
	Address       string `json:"address"`
	Image         string `json:"image"`
	PaslonID      int    `json:"paslon_id"`
}
