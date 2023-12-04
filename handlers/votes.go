package handlers

import (
	"net/http"

	dto "github.com/Devazt/go-restapi-gin/dto/results"
	votesdto "github.com/Devazt/go-restapi-gin/dto/votes"
	"github.com/Devazt/go-restapi-gin/models"
	"github.com/Devazt/go-restapi-gin/repositories"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type voteHandler struct {
	VoteRepo repositories.VoteRepo
}

func VoteHandler(voteRepo repositories.VoteRepo) *voteHandler {
	return &voteHandler{VoteRepo: voteRepo}
}

func (h *voteHandler) FindVotes(c *gin.Context) {
	votes, err := h.VoteRepo.FindVotes()
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	votesLength := len(votes)

	response := models.VoteResponse{
		Votes:      votes,
		TotalVotes: votesLength,
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: response})
}

func (h *voteHandler) Vote(c *gin.Context) {
	request := new(votesdto.VoteReq)
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	userLogin, _ := c.Get("userLogin")
	UserId := userLogin.(jwt.MapClaims)["id"].(float64)

	votes, err := h.VoteRepo.FindVote(int(UserId))
	if votes {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "You have voted"})
		return
	}

	data := models.Vote{
		UserID:   int(UserId),
		PaslonID: request.PaslonID,
	}

	response, err := h.VoteRepo.Vote(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: response})

}
