package repositories

import (
	"context"
	"database/sql"
	"errors"
	"hex/internal/core/domain"
)

type userRepositoryPostgres struct {
	db *sql.DB
}

func NewUserRepositoryPostgres(db *sql.DB) domain.UserRepository {
	return &userRepositoryPostgres{
		db: db,
	}
}

func (u userRepositoryPostgres) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	createdUser := domain.User{
		ID:   "123",
		Name: user.Name,
	}

	return createdUser, nil
}

func (u userRepositoryPostgres) GetUserByID(ctx context.Context, userID string) (domain.User, error) {
	if userID == "123" {
		user := domain.User{
			ID:   "123",
			Name: "John Doe",
		}
		return user, nil
	}

	return domain.User{}, errors.New("user not found")
}
