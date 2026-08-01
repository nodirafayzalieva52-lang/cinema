package handlers

import (
	"github.com/nodirafayzalieva52-lang/cinema/api-gateway/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
