package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	userpb "github.com/nodirafayzalieva52-lang/cinema/user-service/userpb"
)

type UserResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int32  `json:"age"`
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Age      int32  `json:"age"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	ID    int64  `json:"id" binding:"required"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int32  `json:"age"`
}

type RespErr struct {
	Error string `json:"error"`
}

// GetUser
//
// @Summary Get User
// @Description Get user by id
// @Tags User
// @Security BearerAuth
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} UserResponse
// @Failure 400 {object} RespErr
// @Failure 500 {object} RespErr
// @Router /api/user/get/{user_id} [get]
func (h *handler) GetUser(c *gin.Context) {
	idStr := c.Param("user_id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, RespErr{Error: "user_id is required"})
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, RespErr{Error: "invalid user_id format"})
		return
	}

	response, err := h.serviceManager.UserService().GetByID(c.Request.Context(), &userpb.GetUserRequest{
		Id: id,
	})
	if err != nil {
		log.Println("handler GetUser error:", err)
		c.JSON(http.StatusInternalServerError, RespErr{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		ID:    response.GetId(),
		Name:  response.GetName(),
		Age:   int32(response.GetAge()),
		Email: response.GetEmail(),
	})
}

// CreateUser
//
// @Summary Create User
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param body body CreateUserRequest true "User Data"
// @Success 201 {object} UserResponse
// @Failure 400 {object} RespErr
// @Failure 500 {object} RespErr
// @Router /api/user/create [post]
func (h *handler) CreateUser(c *gin.Context) {
	var body CreateUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, RespErr{Error: "invalid request body"})
		return
	}

	response, err := h.serviceManager.UserService().CreateUser(c.Request.Context(), &userpb.CreateUserRequest{
		Name:     body.Name,
		Email:    body.Email,
		Age:      int32(body.Age),
		Password: body.Password,
	})
	if err != nil {
		log.Println("handler CreateUser error:", err)
		c.JSON(http.StatusInternalServerError, RespErr{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, UserResponse{
		ID:    response.GetId(),
		Name:  response.GetName(),
		Email: response.GetEmail(),
		Age:   int32(response.GetAge()),
	})
}

// UpdateUser
//
// @Summary Update User
// @Description Update existing user info
// @Tags User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body UpdateUserRequest true "User Update Data"
// @Success 200 {object} UserResponse
// @Failure 400 {object} RespErr
// @Failure 500 {object} RespErr
// @Router /api/user/update [put]
func (h *handler) UpdateUser(c *gin.Context) {
	var body UpdateUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, RespErr{Error: "invalid request body"})
		return
	}

	response, err := h.serviceManager.UserService().UpdateUser(c.Request.Context(), &userpb.UpdateUserRequest{
		Id:    body.ID,
		Name:  body.Name,
		Email: body.Email,
		Age:   int32(body.Age),
	})
	if err != nil {
		log.Println("handler UpdateUser error:", err)
		c.JSON(http.StatusInternalServerError, RespErr{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		ID:    response.GetId(),
		Name:  response.GetName(),
		Email: response.GetEmail(),
		Age:   int32(response.GetAge()),
	})
}

// DeleteUser
//
// @Summary Delete User
// @Description Delete user by id
// @Tags User
// @Security BearerAuth
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} RespErr
// @Failure 400 {object} RespErr
// @Failure 500 {object} RespErr
// @Router /api/user/delete/{user_id} [delete]
func (h *handler) DeleteUser(c *gin.Context) {
	idStr := c.Param("user_id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, RespErr{Error: "invalid user_id format"})
		return
	}

	// Вызов gRPC DeleteUser вместо Delete
	_, err = h.serviceManager.UserService().DeleteUser(c.Request.Context(), &userpb.DeleteUserRequest{
		Id: id,
	})
	if err != nil {
		log.Println("handler DeleteUser error:", err)
		c.JSON(http.StatusInternalServerError, RespErr{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}