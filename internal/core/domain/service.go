package domain

import "context"

// UserService defines the contract for the user-related operations in the domain layer.
// Implementation of the UserService will reside in the application layer
type UserService interface {
	CreateUser(ctx context.Context, user User) (User, error)
	GetUser(ctx context.Context, userID string) (User, error)
}
