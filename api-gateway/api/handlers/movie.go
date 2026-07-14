package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	moviepb "github.com/nodirafayzalieva52-lang/cinema/movie-service/movie"
	"github.com/nodirafayzalieva52-lang/cinema/api-gateway/models"
)

func (h *handler) CreateMovie(c *gin.Context) {
	var body models.CreateMovieRequest

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	response, err := h.serviceManager.MovieService().Create(
		c.Request.Context(),
		&moviepb.CreateMovieRequest{
			Title:       body.Title,
			Description: body.Description,
		})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}
