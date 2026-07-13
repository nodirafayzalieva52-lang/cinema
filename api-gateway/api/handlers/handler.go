package handlers

import "api-gateway/services"

type handler struct {
	serviceManager services.IServiceManager
}

func NewHandler(serviceManager services.IServiceManager) *handler {
	return &handler{
		serviceManager: serviceManager,
	}
}
