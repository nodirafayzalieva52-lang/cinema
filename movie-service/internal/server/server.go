package server

import (
	"context"
	"errors"
	movie "nd/movie"
	"nd/internal/models"
	"nd/internal/service"
	"nd/pkg/logger"
)

type Server struct {
	movie.UnimplementedMovieServiceServer
	lg      logger.Logger
	service service.Service
}

func New(lg logger.Logger, service service.Service) *Server {
	return &Server{
		lg:      lg,
		service: service,
	}
}

func (s *Server) Add(ctx context.Context, req *movie.CreateMovieRequest) (*movie.CreateMovieResponse, error) {

	request := models.Movie{
		Title: req.Title,
		Description: req.Description,
	}

	movieID, err := s.service.Add(ctx, request)
	if err != nil {
		return nil, err
	}

	return &movie.CreateMovieResponse{
		Id: int64(movieID),
	}, nil
}

func (s *Server) GetByID(ctx context.Context, req *movie.GetMovieRequest) (*movie.GetMovieResponse, error) {
	id := req.Id
	if id < 1 {
		return &movie.GetMovieResponse{}, errors.New("invalid id")
	}

	MovieFromDB, err := s.service.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &movie.GetMovieResponse{
		Id:    int64(MovieFromDB.ID),
		Title: MovieFromDB.Title,
		Description: MovieFromDB.Description,
	}, nil
}

func (s *Server) UpdateMovie(ctx context.Context, req *movie.UpdateMovieRequest) (*movie.UpdateMovieResponse, error) {
	request := models.UpdateMovie{
		ID:    int(req.Id),
		Title: req.Title,
		Description: req.Description,
	}

	err := s.service.Update(ctx, request)
	if err != nil {
		return nil, err
	}

	return &movie.UpdateMovieResponse{
		Code:    0,
		Message: "movie successful updated",
	}, nil
}
