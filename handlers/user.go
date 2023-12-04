package handlers

import (
	"net/http"
	"strconv"
	"time"

	dto "github.com/Devazt/go-restapi-gin/dto/results"
	usersdto "github.com/Devazt/go-restapi-gin/dto/users"
	"github.com/Devazt/go-restapi-gin/models"
	"github.com/Devazt/go-restapi-gin/pkg/bcrypt"
	"github.com/Devazt/go-restapi-gin/repositories"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	UserRepo repositories.UserRepo
}

func UserHandler(userRepo repositories.UserRepo) *userHandler {
	return &userHandler{UserRepo: userRepo}
}

func (h *userHandler) FindUsers(c *gin.Context) {
	users, err := h.UserRepo.FindUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: users})
}

func (h *userHandler) FindUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := h.UserRepo.FindUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convUserRes(user)})
}

func (h *userHandler) CreateUser(c *gin.Context) {
	request := usersdto.CreateUserReq{}
	if err := c.ShouldBindJSON(&request); err != nil {
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
		Password:  string(password),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	response, err := h.UserRepo.CreateUser(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: response})
}

func convUserRes(user models.User) usersdto.UserRes {
	return usersdto.UserRes{
		ID:   user.ID,
		Name: user.Name,
	}
}
