package controller

import (
	"Facopi-Vol2/Models"
	KeuanganFacopiModels "Facopi-Vol2/Models/KeuanganFacopiModels"
	KeuanganFacopiRepository "Facopi-Vol2/Repository/KeuanganFacopiRepository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetKeuanganFacopiByID(c *gin.Context) {
	var request KeuanganFacopiModels.KeuanganFakopi
	var response Models.BaseResponseModels

	response = KeuanganFacopiRepository.GetKeuanganFacopiByID(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func InsertKeuanganFacopi(c *gin.Context) {
	var request KeuanganFacopiModels.KeuanganFakopi
	var response Models.BaseResponseModels

	if err := c.ShouldBindJSON(&request); err != nil {
		response = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = KeuanganFacopiRepository.InsertKeuanganFacopi(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateKeuanganFacopi(c *gin.Context) {
	var request KeuanganFacopiModels.KeuanganFakopi
	var response Models.BaseResponseModels

	if err := c.ShouldBindJSON(&request); err != nil {
		response = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = KeuanganFacopiRepository.UpdateKeuanganFacopi(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func DeleteKeuanganFacopi(c *gin.Context) {
	var request KeuanganFacopiModels.KeuanganFakopi
	var response Models.BaseResponseModels

	if err := c.ShouldBindJSON(&request); err != nil {
		response = Models.BaseResponseModels{
			CodeResponse:  400,
			HeaderMessage: "Bad Request",
			Message:       err.Error(),
			Data:          nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response = KeuanganFacopiRepository.DeleteKeuanganFacopi(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

