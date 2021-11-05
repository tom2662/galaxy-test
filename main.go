package main

import "galaxy/repository"

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

	router.Run(":8000") //running application, Default port is 8080
}
