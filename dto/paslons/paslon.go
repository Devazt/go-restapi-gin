package paslonsdto

type CreatePaslonReq struct {
	Name          string `json:"name" form:"name" validate:"required"`
	VisionMission string `json:"vision_mission" form:"vision_mission" validate:"required"`
	Image         string `json:"image" form:"image" validate:"required"`
}

type UpdatePaslonReq struct {
	Name          string `json:"name" form:"name"`
	VisionMission string `json:"vision_mission" form:"vision_mission"`
	Image         string `json:"image" form:"image"`
}

type PaslonRes struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Serial        int    `json:"serial"`
	VisionMission string `json:"vision_mission"`
	Image         string `json:"image"`
}
