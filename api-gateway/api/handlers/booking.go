package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	bookingpb "github.com/nodirafayzalieva52-lang/cinema/booking-service/bookingpb"
)

type BookingResponse struct {
	ID        int64  `json:"id"`
	UserID    int64  `json:"user_id"`
	MovieID   int64  `json:"movie_id"`
	Status    string `json:"status"`
}

type CreateBookingRequest struct {
	UserID     int64 `json:"user_id"`
	MovieID    int64 `json:"movie_id"`
}

// GetBooking
//
// @Summary Get Booking
// @Description Get booking by ID
// @Tags Booking
// @Security BearerAuth
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {object} BookingResponse
// @Failure 400 {object} RespErr
// @Failure 500 {object} RespErr
// @Router /api/booking/get/{id} [get]
func (h *handler) GetBooking(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, RespErr{Error: "id is required"})
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, RespErr{Error: "invalid booking id format"})
		return
	}

	response, err := h.serviceManager.BookingService().GetBooking(c.Request.Context(), &bookingpb.GetBookingRequest{
		Id: int32(id),
	})
	if err != nil {
		log.Println("handler GetBooking error:", err)
		c.JSON(http.StatusInternalServerError, RespErr{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, BookingResponse{
		ID: int64(response.Booking.Id),
		UserID: int64(response.Booking.UserId),
		MovieID: int64(response.Booking.MovieId),
		Status: response.Booking.GetStatus().String(),
	})
}

// GetUserBookings
//
// @Summary Get User Bookings
// @Description Get all bookings for a specific user
// @Tags Booking
// @Security BearerAuth
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {array} BookingResponse
// @Failure 400 {object} RespErr
// @Failure 500 {object} RespErr
// @Router /api/booking/user/{user_id} [get]
func (h *handler) GetUserBookings(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, RespErr{Error: "invalid user_id format"})
		return
	}

	response, err := h.serviceManager.BookingService().GetUserBookings(c.Request.Context(), &bookingpb.GetUserBookingsRequest{
		UserId: int32(userID),
	})
	if err != nil {
		log.Println("handler GetUserBookings error:", err)
		c.JSON(http.StatusInternalServerError, RespErr{Error: err.Error()})
		return
	}

	var bookings []BookingResponse
	for _, b := range response.GetBookings() {
		bookings = append(bookings, BookingResponse{
			ID: int64(b.Id),
			UserID: int64(b.UserId),
			MovieID: int64(b.MovieId),
			Status: b.GetStatus().String(),
		})
	}

	c.JSON(http.StatusOK, bookings)
}

// CreateBooking
//
// @Summary Create Booking
// @Description Create a new booking
// @Tags Booking
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body CreateBookingRequest true "Booking Data"
// @Success 201 {object} BookingResponse
// @Failure 400 {object} RespErr
// @Failure 500 {object} RespErr
// @Router /api/booking/create [post]
func (h *handler) CreateBooking(c *gin.Context) {
	var body CreateBookingRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, RespErr{Error: "invalid request body"})
		return
	}

	response, err := h.serviceManager.BookingService().CreateBooking(c.Request.Context(), &bookingpb.CreateBookingRequest{
		UserId:  int32(body.UserID),
		MovieId: int32(body.MovieID),
	})
	if err != nil {
		log.Println("handler CreateBooking error:", err)
		c.JSON(http.StatusInternalServerError, RespErr{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, BookingResponse{
		ID: int64(response.Booking.Id),
		UserID: int64(response.Booking.UserId),
		MovieID: int64(response.Booking.MovieId),
		Status: response.Booking.GetStatus().String(),
	})
}

// CancelBooking
//
// @Summary Cancel Booking
// @Description Cancel/Delete booking by ID
// @Tags Booking
// @Security BearerAuth
// @Produce json
// @Param id path int true "Booking ID"
// @Success 200 {object} RespErr
// @Failure 400 {object} RespErr
// @Failure 500 {object} RespErr
// @Router /api/booking/cancel/{id} [delete]
func (h *handler) CancelBooking(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, RespErr{Error: "invalid booking id format"})
		return
	}

	_, err = h.serviceManager.BookingService().CancelBooking(c.Request.Context(), &bookingpb.CancelBookingRequest{
		Id: int32(id),
	})
	if err != nil {
		log.Println("handler CancelBooking error:", err)
		c.JSON(http.StatusInternalServerError, RespErr{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "booking cancelled successfully"})
}