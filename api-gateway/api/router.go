package api

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/nodirafayzalieva52-lang/cinema/api-gateway/docs"
	"github.com/nodirafayzalieva52-lang/cinema/api-gateway/pkg/rabbitmq"

	"github.com/nodirafayzalieva52-lang/cinema/api-gateway/api/handlers"
	"github.com/nodirafayzalieva52-lang/cinema/api-gateway/config"
	"github.com/nodirafayzalieva52-lang/cinema/api-gateway/services"
)

type Option struct {
	Conf           config.Config
	ServiceManager services.IServiceManager
	RabbitClient   *rabbitmq.Rabbit
}

func New(option Option) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())

	handler := handlers.NewHandler(
		option.ServiceManager,
		option.RabbitClient,
	)

	api := router.Group("/api")

	api.POST("/user/get", handler.GetUser)

	api.POST("/movie/create", handler.CreateMovie)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}