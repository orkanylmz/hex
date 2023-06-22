package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"hex/internal/core/domain"
	v0 "hex/internal/grpc/pb/v0"
)

func (u *UserServerV0) CreateUser(ctx context.Context, req *v0.CreateUserRequest) (*v0.CreateUserResponse, error) {
	user := domain.User{
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}

	createdUser, err := u.svc.CreateUser(ctx, user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v0.CreateUserResponse{User: &v0.User{
		Id:    createdUser.ID,
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}}, nil

}
