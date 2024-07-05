package main

import (
	"log"

	rgm "github.com/egorus1442/Report-Generation-Microservice"
	"github.com/egorus1442/Report-Generation-Microservice/pkg/handler"
	"github.com/egorus1442/Report-Generation-Microservice/pkg/repository"
	"github.com/egorus1442/Report-Generation-Microservice/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(rgm.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
