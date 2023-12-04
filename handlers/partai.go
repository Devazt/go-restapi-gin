package handlers

import (
	"net/http"
	"strconv"
	"time"

	partaisdto "github.com/Devazt/go-restapi-gin/dto/partais"
	dto "github.com/Devazt/go-restapi-gin/dto/results"
	"github.com/Devazt/go-restapi-gin/models"
	"github.com/Devazt/go-restapi-gin/repositories"
	"github.com/gin-gonic/gin"
)

type partaiHandler struct {
	PartaiRepo repositories.PartaiRepo
}

func PartaiHandler(partaiRepo repositories.PartaiRepo) *partaiHandler {
	return &partaiHandler{PartaiRepo: partaiRepo}
}

func (p *partaiHandler) FindPartais(c *gin.Context) {
	partais, err := p.PartaiRepo.FindPartais()
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: partais})
}

func (p *partaiHandler) FindPartai(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	partai, err := p.PartaiRepo.FindPartai(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: convPartaiRes(partai)})
}

func (p *partaiHandler) CreatePartai(c *gin.Context) {
	dataFile, _ := c.Get("dataFile")

	request := partaisdto.CreatePartaiReq{
		Name:          c.Request.FormValue("name"),
		Leader:        c.Request.FormValue("leader"),
		VisionMission: c.Request.FormValue("vision_mission"),
		Address:       c.Request.FormValue("address"),
		PaslonID:      c.Request.FormValue("paslon_id"),
	}

	paslonId, err := strconv.Atoi(request.PaslonID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	data := models.Partai{
		Name:          request.Name,
		Leader:        request.Leader,
		VisionMission: request.VisionMission,
		Address:       request.Address,
		PaslonID:      paslonId,
		Image:         dataFile.(string),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	response, err := p.PartaiRepo.CreatePartai(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: response})
}

func (p *partaiHandler) UpdatePartai(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	partai, err := p.PartaiRepo.FindPartai(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	dataFile, _ := c.Get("dataFile")

	request := partaisdto.UpdatePartaiReq{
		Name:          c.Request.FormValue("name"),
		Leader:        c.Request.FormValue("leader"),
		VisionMission: c.Request.FormValue("vision_mission"),
		Address:       c.Request.FormValue("address"),
		PaslonID:      c.Request.FormValue("paslon_id"),
	}

	paslonId, err := strconv.Atoi(request.PaslonID)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	if request.Name != "" {
		partai.Name = request.Name
	}

	if request.Leader != "" {
		partai.Leader = request.Leader
	}

	if request.VisionMission != "" {
		partai.VisionMission = request.VisionMission
	}

	if request.Address != "" {
		partai.Address = request.Address
	}

	if dataFile != "" {
		partai.Image = dataFile.(string)
	}

	if request.PaslonID != "0" {
		partai.PaslonID = paslonId
	}

	partai.UpdatedAt = time.Now()

	response, err := p.PartaiRepo.UpdatePartai(partai)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: response})
}

func (p *partaiHandler) DeletePartai(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	partai, err := p.PartaiRepo.FindPartai(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
		return
	}

	response, err := p.PartaiRepo.DeletePartai(partai, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: response})

}

func convPartaiRes(partai models.Partai) partaisdto.PartaiRes {
	return partaisdto.PartaiRes{
		ID:            partai.Id,
		Name:          partai.Name,
		Leader:        partai.Leader,
		Serial:        partai.Serial,
		VisionMission: partai.VisionMission,
		Address:       partai.Address,
	}
}
