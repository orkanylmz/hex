package grpc

import (
	"database/sql"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"hex/internal/core/application"
	"hex/internal/core/ports"
	v0 "hex/internal/grpc/pb/v0"
	"hex/internal/infrastructure/repositories"
	"log"
	"net"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(db *sql.DB) *Adapter {
	// Initialize and configure any dependencies or connections for gRPC communication
	return &Adapter{db: db}
}

type UserServerV0 struct {
	v0.UnimplementedUserServiceServer
	svc ports.UserServiceServer
}

func NewUserServiceServer(userService ports.UserServiceServer) *UserServerV0 {
	return &UserServerV0{
		svc: userService,
	}
}

// ServeGRPC registers the gRPC server and starts listening for requests
func (a *Adapter) ServeGRPC() error {
	// Initialize and configure the gRPC server

	listener, err := net.Listen("tcp", "5000")

	if err != nil {
		return err
	}

	s := grpc.NewServer()

	userRepository := repositories.NewUserRepositoryPostgres(a.db)
	userService := application.NewUserService(userRepository)

	// Instantiate the UserService from the core domain
	srv := NewUserServiceServer(userService)
	v0.RegisterUserServiceServer(s, srv)

	// Start listening for gRPC requests
	reflection.Register(s)

	go func(s *grpc.Server, lis *net.Listener) {
		if err := s.Serve(*lis); err != nil {
			log.Fatal(fmt.Sprintf("Failed to serve gRPC: %v", err))
		}
	}(s, &listener)

	fmt.Println("service started successfully! Ready to accept incoming requests...")

	return nil
}
