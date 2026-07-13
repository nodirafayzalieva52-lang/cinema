package movie

import (
	"context"
	"fmt"
	movie "movie/movie"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
) 

type MovieGateway interface {
	GetMovie(ctx context.Context, id int) (*movie.GetMovieResponse, error) 
}

type gateway struct {
	client movie.MovieServerClient
}

func New(address string) (MovieGateway, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to movie service: %w", err)
	}

	return &gateway{
		client: movie.NewMovieServiceClient(conn),
	}, nil
}

func (g *gateway) GetMovie(ctx context.Context, id int64) (*movie.GetMovieResponse, error) {
	return g.client.GetByID(ctx, &movie.GetMovieRequest{Id: id})
}
