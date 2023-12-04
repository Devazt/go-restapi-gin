package handlers

import (
	"fmt"
	"net/http"
	"time"

	authsdto "github.com/Devazt/go-restapi-gin/dto/auth"
	dto "github.com/Devazt/go-restapi-gin/dto/results"
	"github.com/Devazt/go-restapi-gin/models"
	"github.com/Devazt/go-restapi-gin/pkg/bcrypt"
	jwtToken "github.com/Devazt/go-restapi-gin/pkg/jwt"
	"github.com/Devazt/go-restapi-gin/repositories"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type authHandler struct {
	AuthRepo repositories.AuthRepo
}

func AuthHandler(authRepo repositories.AuthRepo) *authHandler {
	return &authHandler{AuthRepo: authRepo}
}

func (h *authHandler) Register(c *gin.Context) {
	request := new(authsdto.AuthReq)

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	password, err := bcrypt.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	data := models.User{
		Name:      request.Name,
		Address:   request.Address,
		Gender:    request.Gender,
		Username:  request.Username,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	response, err := h.AuthRepo.Register(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: response})
}

func (h *authHandler) Login(c *gin.Context) {
	request := new(authsdto.LoginReq)

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	user := models.User{
		Username: request.Username,
		Password: request.Password,
	}

	user, err := h.AuthRepo.Login(user.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	isValid := bcrypt.CheckPasswordHash(request.Password, user.Password)
	if !isValid {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "wrong email or password"})
	}

	claims := jwt.MapClaims{}
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		fmt.Println(errGenerateToken)
		return
	}

	loginResponse := authsdto.LoginRes{
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
		Gender:   user.Gender,
		Token:    token,
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: loginResponse})
}
