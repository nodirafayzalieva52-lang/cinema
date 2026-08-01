package handler

import (
	"github.com/nodirafayzalieva52-lang/cinema/booking-service/pkg/rabbitmq"
	"context"
)

type Handler struct {
	rabbitmqManager *rabbitmq.Manager

}

func New(rabbitmqManager *rabbitmq.Manager) Handler {
	return Handler{
		rabbitmqManager: rabbitmqManager,
	}
}


func (h *Handler) Start(ctx context.Context) {
	h.rabbitmqManager.Register("create.booking", h.CreateBooking)

	go h.rabbitmqManager.Start(context.Background())
}


func (h *Handler) CreateBooking(ctx context.Context,body []byte) error {
	return nil
}