package grpc

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v0 "hex/internal/grpc/pb/v0"
)

func (u *UserServerV0) GetUser(ctx context.Context, req *v0.GetUserRequest) (*v0.GetUserResponse, error) {
	foundUser, err := u.svc.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &v0.GetUserResponse{User: &v0.User{
		Id:    foundUser.ID,
		Name:  foundUser.Name,
		Email: foundUser.Email,
	}}, nil
}
