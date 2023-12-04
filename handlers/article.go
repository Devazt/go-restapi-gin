package handlers

import (
	"net/http"
	"strconv"
	"time"

	articlesdto "github.com/Devazt/go-restapi-gin/dto/articles"
	dto "github.com/Devazt/go-restapi-gin/dto/results"
	"github.com/Devazt/go-restapi-gin/models"
	"github.com/Devazt/go-restapi-gin/repositories"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
)

type articleHandler struct {
	ArticleRepo repositories.ArticleRepo
}

func ArticleHandler(articleRepo repositories.ArticleRepo) *articleHandler {
	return &articleHandler{articleRepo}
}

func (h *articleHandler) FindArticles(c *gin.Context) {
	articles, err := h.ArticleRepo.FindArticles()
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: articles})
}

func (h *articleHandler) FindArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := h.ArticleRepo.FindArticle(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convArticleRes(article)})

}

func (h *articleHandler) CreateArticle(c *gin.Context) {
	dataFile, _ := c.Get("dataFile")

	request := articlesdto.CreateArticleReq{
		Title:   c.Request.FormValue("title"),
		Content: c.Request.FormValue("content"),
	}

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	userLogin := c.MustGet("userLogin")
	UserId := userLogin.(jwt.MapClaims)["id"].(float64)

	data := models.Article{
		Title:     request.Title,
		Content:   request.Content,
		Image:     dataFile.(string),
		UserID:    int(UserId),
		CreatedAt: time.Now(),
	}

	response, err := h.ArticleRepo.CreateArticle(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: response})
}

func (h *articleHandler) UpdateArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := h.ArticleRepo.FindArticle(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	dataFile, _ := c.Get("dataFile")

	userLogin, _ := c.Get("userLogin")
	UserId := userLogin.(jwt.MapClaims)["id"].(float64)

	request := articlesdto.UpdateArticleReq{
		Title:   c.Request.FormValue("title"),
		Content: c.Request.FormValue("content"),
		UserID:  int(UserId),
	}

	if request.Title != "" {
		article.Title = request.Title
	}
	if request.Content != "" {
		article.Content = request.Content
	}
	if dataFile != "" {
		article.Image = dataFile.(string)
	}
	if request.UserID != 0 {
		article.UserID = int(UserId)
	}

	article.UpdatedAt = time.Now()

	response, err := h.ArticleRepo.UpdateArticle(article)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: response})
}

func (h *articleHandler) DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := h.ArticleRepo.FindArticle(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	response, err := h.ArticleRepo.DeleteArticle(article, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: response})
}

func convArticleRes(c models.Article) articlesdto.ArticleRes {
	return articlesdto.ArticleRes{
		ID:      c.ID,
		Title:   c.Title,
		Image:   c.Image,
		Content: c.Content,
		UserID:  c.UserID,
	}
}
