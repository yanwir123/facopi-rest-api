package controller

import (
	"Facopi-Vol2/Models"
	DataFacopiModels "Facopi-Vol2/Models/DataFacopiModels"
	DataFacopiRepository "Facopi-Vol2/Repository/DataFacopiRepository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDataFacopiByID(c *gin.Context) {
	var request DataFacopiModels.DataFacopi
	var response Models.BaseResponseModels

	response = DataFacopiRepository.GetDataFacopiByID(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func InsertDataFacopi(c *gin.Context) {
	var request DataFacopiModels.DataFacopi
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

	response = DataFacopiRepository.InsertDataFacopi(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func UpdateDataFacopi(c *gin.Context) {
	var request DataFacopiModels.DataFacopi
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

	response = DataFacopiRepository.UpdateDataFacopi(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func DeleteDataFacopi(c *gin.Context) {
	var request DataFacopiModels.DataFacopi
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

	response = DataFacopiRepository.DeleteDataFacopi(request)
	if response.CodeResponse != 200 {
		c.JSON(response.CodeResponse, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

