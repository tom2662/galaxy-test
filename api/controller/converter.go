package controller

import (
	"galaxy/api/service"
	"galaxy/models"
	"galaxy/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConverterController struct {
	service service.ConverterService
}

func NewConverterController(s service.ConverterService) ConverterController {
	return ConverterController{
		service: s,
	}
}

func (p *ConverterController) ConvertRoman(ctx *gin.Context) {

	var converter models.Converter

	if err := ctx.ShouldBindJSON(&converter); err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Error Bind data")
		return
	}

	if converter.Number == "" {
		util.ErrorJSON(ctx, http.StatusBadRequest, "Number is required")
		return
	}

	response, err := p.service.ConvertRoman(converter)
	if err != nil {
		util.ErrorJSON(ctx, http.StatusBadRequest, "I have no idea what you are talking about")
		return
	}

	ctx.JSON(http.StatusOK, &util.Response{
		Success: true,
		Message: "Conversion Result",
		Data:    &response})
}
