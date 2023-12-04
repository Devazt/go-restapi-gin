package router

import "github.com/gin-gonic/gin"

func RouteInit(g *gin.RouterGroup) {

	UserRoutes(g)
	AuthRoutes(g)
	PaslonRoutes(g)
	PartaiRoutes(g)
	ArticleRoutes(g)
	VoteRoutes(g)

}
