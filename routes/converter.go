package routes

import (
	"server/api/controller"
	"server/infrastructure"
)

type ConverterRoute struct {
	Controller controller.ConverterController
	Handler    base.GinRouter
}

func NewConverterRoute(
	controller controller.ConverterController,
	handler infrastructure.GinRouter,

) ConverterRoute {
	return ConverterRoute{
		Controller: controller,
		Handler:    handler,
	}
}

func (p ConverterRoute) Setup() {
	post := p.Handler.Gin.Group("/converter")
	{
		converter.POST("/", p.Controller.ConvertNum)
	}
}
