package handlers

import (
	"github.com/nodirafayzalieva52-lang/cinema/api-gateway/config"
	"github.com/nodirafayzalieva52-lang/cinema/api-gateway/services"
)

type HandlerConfig struct {
	ServiceManager services.IServiceManager
	Config         config.Config
}

type handler struct {
	serviceManager services.IServiceManager
	cfg            config.Config
}

func New(c *HandlerConfig) *handler {
	return &handler{
		serviceManager: c.ServiceManager,
		cfg:            c.Config,
	}
}