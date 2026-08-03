package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	moviepb "github.com/nodirafayzalieva52-lang/cinema/movie-service/movie"
)

type MovieResponse struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
}

type CreateMovieRequest struct {
	Title       string  `json:"title"`
}

type UpdateMovieRequest struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
}

// GetMovie
//
// @Summary Get Movie
// @Description Get movie by ID
// @Tags Movie
// @Security BearerAuth
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} MovieResponse
// @Failure 400 {object} RespErr
// @Failure 500 {object} RespErr
// @Router /api/movie/get/{id} [get]
func (h *handler) GetMovie(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, RespErr{Error: "id is required"})
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, RespErr{Error: "invalid movie id format"})
		return
	}

	response, err := h.serviceManager.MovieService().GetByID(c.Request.Context(), &moviepb.GetMovieRequest{
		Id: id,
	})
	if err != nil {
		log.Println("handler GetMovie error:", err)
		c.JSON(http.StatusInternalServerError, RespErr{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, MovieResponse{
		ID:          response.GetId(),
		Title:       response.GetTitle(),
	})
}

// CreateMovie
//
// @Summary Create Movie
// @Description Create a new movie entry
// @Tags Movie
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body CreateMovieRequest true "Movie Data"
// @Success 201 {object} MovieResponse
// @Failure 400 {object} RespErr
// @Failure 500 {object} RespErr
// @Router /api/movie/create [post]
func (h *handler) CreateMovie(c *gin.Context) {
	var body CreateMovieRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, RespErr{Error: "invalid request body"})
		return
	}

	response, err := h.serviceManager.MovieService().Create(c.Request.Context(), &moviepb.CreateMovieRequest{
		Title:       body.Title,
	})
	if err != nil {
		log.Println("handler CreateMovie error:", err)
		c.JSON(http.StatusInternalServerError, RespErr{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, MovieResponse{
		ID:          response.GetId(),
	})
}

// UpdateMovie
//
// @Summary Update Movie
// @Description Update movie details
// @Tags Movie
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body UpdateMovieRequest true "Movie Update Data"
// @Success 200 {object} MovieResponse
// @Failure 400 {object} RespErr
// @Failure 500 {object} RespErr
// @Router /api/movie/update [put]
func (h *handler) UpdateMovie(c *gin.Context) {
    var body UpdateMovieRequest
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, RespErr{Error: "invalid request body"})
        return
    }

    response, err := h.serviceManager.MovieService().Update(c.Request.Context(), &moviepb.UpdateMovieRequest{
        Id:    body.ID,
        Title: body.Title,
    })
    if err != nil {
        log.Println("handler UpdateMovie error:", err)
        c.JSON(http.StatusInternalServerError, RespErr{Error: err.Error()})
        return
    }

    // Возвращаем статус из response или обновленный объект
    c.JSON(http.StatusOK, gin.H{
        "code":    response.GetCode(),
        "message": response.GetMessage(),
    })
}

// DeleteMovie
//
// @Summary Delete Movie
// @Description Delete movie by ID
// @Tags Movie
// @Security BearerAuth
// @Produce json
// @Param id path int true "Movie ID"
// @Success 200 {object} RespErr
// @Failure 400 {object} RespErr
// @Failure 500 {object} RespErr
// @Router /api/movie/delete/{id} [delete]
func (h *handler) DeleteMovie(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, RespErr{Error: "invalid movie id format"})
		return
	}

	_, err = h.serviceManager.MovieService().Delete(c.Request.Context(), &moviepb.DeleteMovieRequest{
		Id: id,
	})
	if err != nil {
		log.Println("handler DeleteMovie error:", err)
		c.JSON(http.StatusInternalServerError, RespErr{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "movie deleted successfully"})
}