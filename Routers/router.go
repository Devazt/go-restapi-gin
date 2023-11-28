package router

import "github.com/gin-gonic/gin"

func RouteInit(g *gin.RouterGroup) {

	AuthRoute(g)
	ArticleRoute(g)
	PaslonRoute(g)
	PartaiRoute(g)

}
