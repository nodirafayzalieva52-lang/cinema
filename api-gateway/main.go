package main

import (
	"log"

	"api-gateway/api"
	"api-gateway/config"
	"api-gateway/services"
)

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

	if err := server.Run(conf.HTTPPORT); err != nil {
		log.Fatal("server.Run(): %v", err)
	}
}