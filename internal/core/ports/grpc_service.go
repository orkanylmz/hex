package ports

import (
	"context"
	"hex/internal/core/domain"
)

// Concrete implementation should be in internal/adapters/grpc/server.go

// UserServiceServer defines the methods that the gRPC service should implement
type UserServiceServer interface {
	CreateUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUser(ctx context.Context, id string) (domain.User, error)
}
