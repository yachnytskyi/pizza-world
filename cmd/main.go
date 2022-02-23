package main

import (
	"log"

	"github.com/yachnytskyi/pizza-world/cmd/server"
	"github.com/yachnytskyi/pizza-world/pkg/handler"
	"github.com/yachnytskyi/pizza-world/pkg/repository"
	"github.com/yachnytskyi/pizza-world/pkg/service"
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
