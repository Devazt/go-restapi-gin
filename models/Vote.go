package models

type Vote struct {
	ID       int `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID   int `json:"user_id"`
	PaslonID int `json:"paslon_id"`
}

type VoteUserResponse struct {
	ID       int `json:"id"`
	UserID   int `json:"user_id"`
	PaslonID int `json:"paslon_id"`
}

type VotePaslonResponse struct {
	ID       int `json:"id"`
	PaslonID int `json:"paslon_id"`
}

type VoteResponse struct {
	Votes      []Vote `json:"votes"`
	TotalVotes int    `json:"totalVotes"`
}
