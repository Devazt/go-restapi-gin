package votesdto

type VoteReq struct {
	UserID   int `json:"user_id" form:"user_id" validate:"required"`
	PaslonID int `json:"paslon_id" form:"paslon_id" validate:"required"`
}

type VoteRes struct {
	UserID   int `json:"user_id"`
	PaslonID int `json:"paslon_id"`
}
