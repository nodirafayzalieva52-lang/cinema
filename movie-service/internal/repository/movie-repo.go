package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/nodirafayzalieva52-lang/cinema/movie-service/internal/models"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) Repository { 
	return Repository{db: db}
}

func (r *Repository) Add(ctx context.Context, request models.Movie) (int, error) {
	const query = `
			INSERT INTO mv_movies (title, description, duration, age_limit, created_at)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id`

	var idv int
	err := r.db.QueryRow(ctx, query, request.Title, request.Description, request.Duration, request.Age_Limit, request.CreatedAt).Scan(&idv)
	if err != nil {
		return 0, errors.New("error from r.db.Exec")
	}

	return idv, nil
}

func (r *Repository) GetByID(ctx context.Context, id int64) (models.Movie, error) {
	const query = `SELECT id, title, description, duration, age_limit, created_at FROM mv_movies WHERE id = $1`

	var movie models.Movie
	err := r.db.QueryRow(ctx, query, id).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Duration,
		&movie.Age_Limit,
		&movie.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Movie{}, fmt.Errorf("Movie not found")
		}

		return models.Movie{}, err
	}
	return movie, err

}

func (r *Repository) Update(ctx context.Context, request models.UpdateMovie) error {

	const query = `UPDATE mv_movies SET title = $2, description = $3, duration = $4, age_limit = $5 WHERE id = $1`

	result, err := r.db.Exec(ctx, query, request.ID, request.Title, request.Description, request.Duration, request.Age_Limit)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("Result shoud not be empty")
	}
	return nil
}


func (r *Repository) DeleteMovie(ctx context.Context, id int) error {
	const query = `UPDATE mv_movies SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}