package application

import (
	"context"
	"hex/internal/core/domain"
)

type userService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (s *userService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	// Perform any necessary validation or business rules

	createdUser, err := s.userRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	return createdUser, nil
}

func (s *userService) GetUser(ctx context.Context, userID string) (domain.User, error) {

	user, err := s.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}
