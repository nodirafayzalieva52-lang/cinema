package api

import (
	"github.com/gin-gonic/gin"
	"api-gateway/config"
	"api-gateway/api/handlers"
	"api-gateway/services"
)

type Option struct {
	Conf           config.Config
	ServiceManager services.IServiceManager
}

func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())

	handler := handlers.NewHandler(
		option.ServiceManager,
	)

	api := router.Group("/api")

	api.POST("/user/get", handler.GetUser)

	api.POST("/movie/create", handler.CreateMovie)

	return router
}