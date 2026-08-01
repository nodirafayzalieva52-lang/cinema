package handlers

import (
	"github.com/nodirafayzalieva52-lang/cinema/api-gateway/pkg/rabbitmq"
	"github.com/nodirafayzalieva52-lang/cinema/api-gateway/services"
)

type handler struct {
	serviceManager services.IServiceManager
	rabbitClient *rabbitmq.Rabbit
}

func NewHandler(serviceManager services.IServiceManager, rabbitClient *rabbitmq.Rabbit) *handler {
	return &handler{
		serviceManager: serviceManager,
		rabbitClient: rabbitClient,
	}
}
