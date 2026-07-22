package main

import (
	"log"

	"github.com/nodirafayzalieva52-lang/cinema/api-gateway/api"
	"github.com/nodirafayzalieva52-lang/cinema/api-gateway/config"
	"github.com/nodirafayzalieva52-lang/cinema/api-gateway/services"
)

// @title Cinema Api
// @version v1.0.0
// @description Api for cinema project
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath /
func main() {
	conf, err := config.New("./config/config.env")
	if err != nil {
		log.Fatal(err)
	}

	serviceManager, err := services.NewServiceManager(conf.Services)
	if err != nil {
		log.Fatalf("services.NewServiceManager(): %v", err)
	}

	server := api.New(api.Option{
		ServiceManager: serviceManager,
	})

	if err := server.Run(":8080"); err != nil {
		log.Fatalf("server.Run(): %v", err)
	}
}