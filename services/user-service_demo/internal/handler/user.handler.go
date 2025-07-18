package handler

import (
	"context"
	"fmt"

	"github.com/098765432m/user-service/internal/service"
	"github.com/098765432m/user-service/pb"
)

type UserHanlder struct {
	pb.UnimplementedUserServiceServer
	userService *service.UserServiceImpl
}

func NewUserHanlder(userService *service.UserServiceImpl) *UserHanlder {
	return &UserHanlder{
		userService: userService,
	}
}

func (s *UserHanlder) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.userService.GetUserById(req.Id)
	if err != nil {
		fmt.Println(err)
	}
	return &pb.GetUserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
