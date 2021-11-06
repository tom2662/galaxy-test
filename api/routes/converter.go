package routes

import (
	"galaxy/api/controller"
	"galaxy/base"
)

type ConverterRoute struct {
	Controller controller.ConverterController
	Handler    base.GinRouter
}

func NewConverterRoute(
	controller controller.ConverterController,
	handler base.GinRouter,

) ConverterRoute {
	return ConverterRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (p ConverterRoute) Setup() {
	converter := p.Handler.Gin.Group("/converter")
	{
		converter.POST("/", p.Controller.ConvertRoman)
	}
}
