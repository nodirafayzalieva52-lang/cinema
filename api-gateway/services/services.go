package services

import (
	"fmt"

	"github.com/nodirafayzalieva52-lang/cinema/api-gateway/config"
	bookingpb "github.com/nodirafayzalieva52-lang/cinema/booking-service/bookingpb"
	moviepb "github.com/nodirafayzalieva52-lang/cinema/movie-service/movie"
	userpb "github.com/nodirafayzalieva52-lang/cinema/user-service/userpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	UserService() userpb.UserServiceClient
	MovieService() moviepb.MovieServiceClient
	BookingService() bookingpb.BookingServiceClient
}

type serviceManager struct {
	userService    userpb.UserServiceClient
	movieService   moviepb.MovieServiceClient
	bookingService bookingpb.BookingServiceClient
}

func (s *serviceManager) UserService() userpb.UserServiceClient {
	return s.userService
}

func (s *serviceManager) MovieService() moviepb.MovieServiceClient {
	return s.movieService
}

func (s *serviceManager) BookingService() bookingpb.BookingServiceClient {
	return s.bookingService
}

func NewServiceManager(config config.Services) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns") 

	connUserService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", &config.UserService.Host, config.UserService.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to user service: %w", err)
	}

	connMovieService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", &config.MovieService.Host, config.MovieService.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("failed to connect to movie service: %w", err)
	}

	connBookingService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", &config.BookingService.Host, config.BookingService.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("failed to connect to movie service: %w", err)
	}

	return &serviceManager{
		userService: userpb.NewUserServiceClient(connUserService),
		movieService: moviepb.NewMovieServiceClient(connMovieService),
		bookingService: bookingpb.NewBookingServiceClient(connBookingService),
	}, nil
}