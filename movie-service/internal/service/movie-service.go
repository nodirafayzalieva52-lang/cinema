package service

import (
	"context"
	"fmt"
	"github.com/nodirafayzalieva52-lang/cinema/movie-service/internal/models"
	"github.com/nodirafayzalieva52-lang/cinema/movie-service/internal/repository"
)

type Service struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return Service{repo: repo}
}

func (s *Service) Add(ctx context.Context, request models.Movie) (int, error) {
	err := request.Validate()
	if err != nil {
		return 0, fmt.Errorf("Validation error")
	}

	movieID, err := s.repo.Add(ctx, models.Movie{
		Title: request.Title,
		Description: request.Description,
		Duration: request.Duration,
		Age_Limit: request.Age_Limit,
		CreatedAt: request.CreatedAt,
	})
	if err != nil {
		return 0, fmt.Errorf("error from s.repo.Add")
	}

	return movieID, nil
}

func (s *Service) GetByID(ctx context.Context, id int64) (models.Movie, error) {
	movie, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return models.Movie{}, err
	}
	return movie, nil
}

func (s *Service) Update(ctx context.Context, request models.UpdateMovie) error {

	err := request.Validate()
	if err != nil {
		return fmt.Errorf("Validation error")
	}

	if request.ID < 1 {
		return fmt.Errorf("Validation error")
	}

	err = s.repo.Update(ctx, request)
	if err != nil {
		return err
	}

	return nil
}
