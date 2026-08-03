package api

import (
	"github.com/gin-gonic/gin"
	_ "github.com/nodirafayzalieva52-lang/cinema/api-gateway/docs" // Swagger docs
	"github.com/nodirafayzalieva52-lang/cinema/api-gateway/api/handlers"
	"github.com/nodirafayzalieva52-lang/cinema/api-gateway/config"
	"github.com/nodirafayzalieva52-lang/cinema/api-gateway/services"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Option struct {
	Conf           config.Config
	ServiceManager services.IServiceManager
}

func New(opt Option) *gin.Engine {
	router := gin.Default()

	handler := handlers.New(&handlers.HandlerConfig{
		ServiceManager: opt.ServiceManager,
		Config:         opt.Conf,
	})

	api := router.Group("/api")
	{
		// === USER ROUTES ===
		api.POST("/user/create", handler.CreateUser)
		api.GET("/user/get/:user_id", handler.GetUser)
		api.PUT("/user/update", handler.UpdateUser)
		api.DELETE("/user/delete/:user_id", handler.DeleteUser)

		// === MOVIE ROUTES ===
		api.POST("/movie/create", handler.CreateMovie)
		api.GET("/movie/get/:id", handler.GetMovie)
		api.PUT("/movie/update", handler.UpdateMovie)
		api.DELETE("/movie/delete/:id", handler.DeleteMovie)

		// === BOOKING ROUTES ===
		api.POST("/booking/create", handler.CreateBooking)
		api.GET("/booking/get/:id", handler.GetBooking)
		api.GET("/booking/user/:user_id", handler.GetUserBookings)
		api.DELETE("/booking/cancel/:id", handler.CancelBooking)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}