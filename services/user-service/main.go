package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/MiguelMachado-dev/grpcommerce/ecommerce/proto/user"

	"github.com/MiguelMachado-dev/grpcommerce/services/user-service/config"
	"github.com/MiguelMachado-dev/grpcommerce/services/user-service/handler"
	"github.com/MiguelMachado-dev/grpcommerce/services/user-service/repository"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize repository
	repo, err := repository.NewPostgresRepository(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer repo.Close()

	// Initialize gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create gRPC server
	server := grpc.NewServer()

	// Register our service
	userHandler := handler.NewUserHandler(repo)
	pb.RegisterUserServiceServer(server, userHandler)

	// Register reflection service for gRPC tools
	reflection.Register(server)

	// Start server
	log.Printf("Starting user service on port %d", cfg.Port)
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	log.Printf("Received signal %s, shutting down...", sig.String())

	server.GracefulStop()
}
