package server

import (
	"context"
	"fmt"

	"github.com/098765432m/grpc-micro/user-service/internal/user/service"
	"github.com/098765432m/grpc-micro/user-service/scripts/pb"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	userService *service.UserService
}

func NewUserServer(userService *service.UserService) *UserServer {
	return &UserServer{
		userService: userService,
	}
}

func (us *UserServer) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {

	user, err := us.userService.HandleGetById(req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetUserByIdResponse{
		Id:       user.Id,
		Username: user.Username,
		Password: user.Password,
		FullName: user.FullName,
		Email:    user.Email,
		IsActive: user.IsActive,
	}, nil
}

func (us *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	newUser := service.CreateUser{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
		FullName: req.GetFullName(),
		Email:    req.GetEmail(),
	}

	newId, err := us.userService.HandleCreate(newUser)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserResponse{
		Id: newId,
	}, nil
}

func (us *UserServer) CheckStatus(ctx context.Context, req *pb.CheckStatusRequest) (*pb.CheckStatusResponse, error) {
	isActive, err := us.userService.HandleCheckStatus(req.GetId())
	if err != nil {
		return nil, fmt.Errorf("failed to get user %s: %v", req.GetId(), err)
	}

	return &pb.CheckStatusResponse{
		IsActive: isActive,
	}, nil
}

func (us *UserServer) CheckUserAuthorized(ctx context.Context, req *pb.IsUserAuthorizedRequest) (*pb.IsUserAuthorizedResponse, error) {

	isAuthorized, err := us.userService.HandleCheckUserAuthorized(req.GetId(), req.GetRole())
	if err != nil {
		return nil, err
	}

	return &pb.IsUserAuthorizedResponse{
		IsAuthorized: isAuthorized,
	}, nil
}
