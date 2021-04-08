package main

import (
	http_rest_api_test "github.com/artyomchch/http-rest-api-test"
	"github.com/artyomchch/http-rest-api-test/pkg/handler"
	"github.com/artyomchch/http-rest-api-test/pkg/repository"
	"github.com/artyomchch/http-rest-api-test/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(http_rest_api_test.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
