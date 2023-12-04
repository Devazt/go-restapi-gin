package middlewares

import (
	"errors"
	"net/http"
	"strings"

	dto "github.com/Devazt/go-restapi-gin/dto/results"
	"github.com/Devazt/go-restapi-gin/pkg/jwt"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Auth(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, dto.ErrorResult{Code: http.StatusUnauthorized, Message: errors.New("unauthorized").Error()})
			return
		}

		token = strings.Split(token, " ")[1]
		claims, err := jwt.GetClaims(token)

		if err != nil {
			response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: err.Error()}
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("userLogin", claims)

		next(c)
	}
}
