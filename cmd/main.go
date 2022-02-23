package main

import (
	"log"
	"pizza-world/cmd/server"
	"pizza-world/pkg/handler"
	"pizza-world/pkg/repository"
	"pizza-world/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
