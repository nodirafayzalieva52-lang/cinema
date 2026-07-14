package handlers

import "github.com/nodirafayzalieva52-lang/cinema/api-gateway/services"

type handler struct {
	serviceManager services.IServiceManager
}

func NewHandler(serviceManager services.IServiceManager) *handler {
	return &handler{
		serviceManager: serviceManager,
	}
}
