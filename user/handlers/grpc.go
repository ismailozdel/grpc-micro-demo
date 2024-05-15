package handlers

import (
	"context"
	"log"

	"github.com/ismailozdel/micro/common/proto/user"
	"github.com/ismailozdel/micro/user/types"
	"google.golang.org/grpc"
)

type UserGrpcHandler struct {
	userService types.UserService
	user.UnimplementedUserServiceServer
}

func NewGrpcUsersService(grpc *grpc.Server, userService types.UserService) {
	grpcHandler := &UserGrpcHandler{
		userService: userService,
	}
	user.RegisterUserServiceServer(grpc, grpcHandler)
}

func (s *UserGrpcHandler) GetUser(ctx context.Context, req *user.GetUserRequest) (*user.GetUserResponse, error) {
	users, err := s.userService.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &user.GetUserResponse{User: users}, nil
}

func (s *UserGrpcHandler) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	users, err := s.userService.CreateUser(ctx, &user.User{Name: req.Name, Email: req.Email, Password: req.Password})
	if err != nil {
		return nil, err
	}
	return &user.CreateUserResponse{User: users}, nil
}

func (s *UserGrpcHandler) UpdateUser(ctx context.Context, req *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	data := req.GetUser()
	users, err := s.userService.UpdateUser(ctx, &user.User{Id: data.Id, Name: data.Name, Email: data.Email, Password: data.Password})
	if err != nil {
		return nil, err
	}
	return &user.UpdateUserResponse{User: users}, nil
}

func (s *UserGrpcHandler) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*user.DeleteUserResponse, error) {
	log.Println("Deleting user with id: ", req.Id)
	err := s.userService.DeleteUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &user.DeleteUserResponse{}, nil
}
