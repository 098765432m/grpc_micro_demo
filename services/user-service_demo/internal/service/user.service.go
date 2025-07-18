package service

type User struct {
	ID    string
	Name  string
	Email string
}

type UserService interface {
	GetUserById(id string) (*User, error)
}

type UserServiceImpl struct{}

func (s *UserServiceImpl) GetUserById(id string) (*User, error) {
	return &User{ID: id, Name: "John Smith", Email: "jsmith@as.com"}, nil
}

func (s *UserServiceImpl) CreateAdmin() {

}
