package repositories

import (
	"github.com/Devazt/go-restapi-gin/models"

	"gorm.io/gorm"
)

type VoteRepo interface {
	FindVotes() ([]models.Vote, error)
	FindVote(ID int) (bool, error)
	Vote(vote models.Vote) (models.Vote, error)
}

type voteRepo struct {
	db *gorm.DB
}

func RepoVote(db *gorm.DB) *voteRepo {
	return &voteRepo{db}
}

func (r *voteRepo) FindVotes() ([]models.Vote, error) {
	var votes []models.Vote
	err := r.db.Find(&votes).Error
	return votes, err
}

func (r *voteRepo) FindVote(ID int) (bool, error) {
	var vote bool
	err := r.db.First(&vote, ID).Error
	return vote, err
}

func (r *voteRepo) Vote(vote models.Vote) (models.Vote, error) {
	err := r.db.Create(&vote).Error
	return vote, err
}
