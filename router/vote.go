package router

import (
	"github.com/Devazt/go-restapi-gin/handlers"
	"github.com/Devazt/go-restapi-gin/pkg/middlewares"
	"github.com/Devazt/go-restapi-gin/pkg/postgres"
	"github.com/Devazt/go-restapi-gin/repositories"
	"github.com/gin-gonic/gin"
)

func VoteRoutes(e *gin.RouterGroup) {
	r := repositories.RepoVote(postgres.DB)
	h := handlers.VoteHandler(r)

	e.GET("/votes", h.FindVotes)
	e.POST("/vote", middlewares.Auth(h.Vote))

}
