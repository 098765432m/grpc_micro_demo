package service

import (
	"fmt"

	"github.com/098765432m/grpc-micro/user-service/consts"
	"github.com/098765432m/grpc-micro/user-service/internal/user/domain"
	"github.com/google/uuid"
)

type UserRepo interface {
	FindById(id string) (*domain.User, error)
	Save(u *domain.User) (string, error)
}

type UserService struct {
	repo UserRepo
}

type CreateUser struct {
	Username string
	Password string
	FullName string
	Email    string
}

func (us *UserService) HandleCreate(newUser CreateUser) (string, error) {
	id := uuid.New().String()

	// Hash password
	hashPass := newUser.Password
	u, err := domain.NewUser(id, newUser.Username, hashPass, newUser.FullName, newUser.Email, consts.RoleGuest, false)
	if err != nil {
		return "", fmt.Errorf("failed to create new User: %v", err)
	}

	return us.repo.Save(u)
}

func (us *UserService) HandleGetById(id string) (*domain.User, error) {
	user, err := us.repo.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find user by id %s: %v", id, err)
	}

	return user, nil
}

func (us *UserService) HandleCheckStatus(id string) (bool, error) {
	user, err := us.repo.FindById(id)
	if err != nil {
		return false, fmt.Errorf("failed to find user by id %s: %v", id, err)
	}

	return user.IsActive, nil

}

func (us *UserService) HandleCheckUserAuthorized(id string, role string) (bool, error) {
	user, err := us.repo.FindById(id)
	if err != nil {
		return false, fmt.Errorf("failed to find user by id %s: %v", id, err)
	}

	isAuthorized := user.Role == role

	return isAuthorized, nil
}
