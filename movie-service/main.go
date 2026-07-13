package main

import (
	"context"
	"log"
	"nd/internal/config"
	"nd/internal/repository"
	"nd/internal/server"
	"nd/internal/service"
	"nd/pkg/db"
	"nd/pkg/logger"
	"nd/movie"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	cfg, err := config.New("./config/config.env")
	if err != nil {
		log.Fatal("config.New", err)
	}

	conn, err := db.New(db.Option{
		Host:     cfg.DBHost,
		Port:     cfg.DBPort,
		Movie:     cfg.DBMovie,
		Password: cfg.DBPassword,
		DBName:   cfg.DBName,
	})
	if err != nil {
		log.Fatal("failed to connect to db:", err)
	}
	defer conn.Close()

	lg, err := logger.New(true)
	if err != nil {
		log.Fatal("failed to create logger", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	lis, err := net.Listen(cfg.NETWORK, cfg.ADDRESS)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	movieRepo := repository.New(conn)

	movieService := service.New(movieRepo)

	movieServer := server.New(*lg, movieService)

	movie.RegisterMovieServiceServer(grpcServer, movieServer)

	reflection.Register(grpcServer)

	go func() {
		lg.Info("server listening at %v", zap.String("addr", lis.Addr().String()))
		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	lg.Info("shutting down server...")
	grpcServer.GracefulStop()
	lg.Info("server stopped")
}
