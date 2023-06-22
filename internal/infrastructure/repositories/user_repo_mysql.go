package repositories

import (
	"context"
	"database/sql"
	"errors"
	"hex/internal/core/domain"
)

type userRepositoryMySQL struct {
	db *sql.DB
}

func NewUserRepositoryMySQL(db *sql.DB) domain.UserRepository {
	return &userRepositoryMySQL{
		db: db,
	}
}

func (u userRepositoryMySQL) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	createdUser := domain.User{
		ID:   "123",
		Name: user.Name,
	}

	return createdUser, nil
}

func (u userRepositoryMySQL) GetUserByID(ctx context.Context, userID string) (domain.User, error) {
	if userID == "123" {
		user := domain.User{
			ID:   "123",
			Name: "John Doe",
		}
		return user, nil
	}

	return domain.User{}, errors.New("user not found")
}
