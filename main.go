package main

import (
	"galaxy/api/controller"
	"galaxy/api/repository"
	"galaxy/api/routes"
	"galaxy/api/service"
	"galaxy/base"
)

func init() {
	base.LoadEnv()
}

func main() {

	router := base.NewGinRouter()
	db := base.NewDatabase()

	converterRepository := repository.NewConverterRepository(db)
	converterService := service.NewConverterService(converterRepository)
	converterController := controller.NewConverterController(converterService)
	converterRoute := routes.NewConverterRoute(converterController, router)
	converterRoute.Setup()

	router.Gin.Run(":8000")
}
