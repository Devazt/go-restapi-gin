package handlers

import (
	"net/http"
	"strconv"
	"time"

	paslonsdto "github.com/Devazt/go-restapi-gin/dto/paslons"
	dto "github.com/Devazt/go-restapi-gin/dto/results"
	"github.com/Devazt/go-restapi-gin/models"
	"github.com/Devazt/go-restapi-gin/repositories"
	"github.com/gin-gonic/gin"
)

type paslonHandler struct {
	PaslonRepo repositories.PaslonRepo
}

func PaslonHandler(paslonRepo repositories.PaslonRepo) *paslonHandler {
	return &paslonHandler{PaslonRepo: paslonRepo}
}

func (p *paslonHandler) FindPaslons(c *gin.Context) {
	partais, err := p.PaslonRepo.FindPaslons()
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: partais})
}

func (p *paslonHandler) FindPaslon(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	paslon, err := p.PaslonRepo.FindPaslon(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convPaslonRes(paslon)})

}

func (p *paslonHandler) CreatePaslon(c *gin.Context) {
	dataFile, _ := c.Get("dataFile")

	request := paslonsdto.CreatePaslonReq{
		Name:          c.Request.FormValue("name"),
		VisionMission: c.Request.FormValue("vision_mission"),
	}

	data := models.Paslon{
		Name:          request.Name,
		VisionMission: request.VisionMission,
		Image:         dataFile.(string),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	response, err := p.PaslonRepo.CreatePaslon(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: response})
}

func (p *paslonHandler) UpdatePaslon(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	paslon, err := p.PaslonRepo.FindPaslon(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	dataFile, _ := c.Get("dataFile")
	request := paslonsdto.UpdatePaslonReq{
		Name:          c.Request.FormValue("name"),
		VisionMission: c.Request.FormValue("vision_mission"),
	}

	if request.Name != "" {
		paslon.Name = request.Name
	}

	if request.VisionMission != "" {
		paslon.VisionMission = request.VisionMission
	}

	if dataFile != "" {
		paslon.Image = dataFile.(string)
	}

	paslon.UpdatedAt = time.Now()

	response, err := p.PaslonRepo.UpdatePaslon(paslon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: response})
}

func (p *paslonHandler) DeletePaslon(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	paslon, err := p.PaslonRepo.FindPaslon(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	response, err := p.PaslonRepo.DeletePaslon(paslon, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: response})

}

func convPaslonRes(paslon models.Paslon) paslonsdto.PaslonRes {
	return paslonsdto.PaslonRes{
		ID:            paslon.Id,
		Name:          paslon.Name,
		Serial:        paslon.Serial,
		VisionMission: paslon.VisionMission,
	}
}
