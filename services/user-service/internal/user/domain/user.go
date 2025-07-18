// User Domain Logic

package domain

import (
	"fmt"
)

type User struct {
	Id       string
	Username string
	Password string
	FullName string
	Email    string
	Role     string
	IsActive bool
}

func NewUser(id string, username string, password, fullName string, email string, role string, isActive bool) (*User, error) {
	return &User{
		Id:       id,
		Username: username,
		Password: password,
		FullName: fullName,
		Email:    email,
		Role:     role,
		IsActive: isActive,
	}, nil
}

func (u *User) SetNewpassword(newHashPass string) error {
	if newHashPass == "" {
		return fmt.Errorf("password hash cannot be empty")
	}

	u.Password = newHashPass
	return nil
}

func (u *User) Activate() {
	u.IsActive = true
}

func (u *User) Deactivate() {
	u.IsActive = false
}
