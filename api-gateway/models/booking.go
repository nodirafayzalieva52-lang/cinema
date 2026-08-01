package models

type CreateBookingRequest struct {
	UserID  int64 `json:"userID"`
	MovieID int64 `json:"movieID"`
}
