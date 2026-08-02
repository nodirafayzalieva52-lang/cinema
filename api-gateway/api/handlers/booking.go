package handlers

import (
	"github.com/nodirafayzalieva52-lang/cinema/api-gateway/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateBooking
//
// @Summary Create booking
// @Description Create a new booking (publishes to queue, async processing)
// @Tags Booking
// @security BearerAuth
// @Accept json
// @Produce json
// @Param request body models.CreateBookingRequest true "Booking data"
// @Success 202 {object} handlers.RespOk
// @Failure 400 {object} RespErr
// @Failure 500 {object} RespErr
// @Router /api/booking/create [post]
func (h *handler) CreateBooking(c *gin.Context) {
	var body models.CreateBookingRequest

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	err = h.rabbitClient.Publisher(c.Request.Context(), body, "create.booking")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted,nil)
}
