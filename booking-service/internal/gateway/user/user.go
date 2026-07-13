package user

import (
	"fmt"
	"context"
	user "movie/userpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserGateway interface {
	GetUser(ctx context.Context, id int) (*user.GetUserREsponse, error)
}

type gateway struct {
	client user.UserServiceClient
}

func New(address string) (UserGateway, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to user service: %w", err)
	}

	return &gateway{
		client: user.NewUserServiceClient(conn),
	}, nil
} 

func (g *gateway) GetUser(ctx context.Context, id int) (*user.GetUserResponse, error) {
	return g.client.GetByID(ctx, &user.GetUserRequest{Id: int64(id)})
}