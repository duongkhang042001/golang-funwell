package service

import "core/internal/domain"

type userService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) domain.UserService {
	return &userService{
		userRepository: userRepo,
	}
}

// CreateUser implements domain.UserService.
func (u *userService) CreateUser(user domain.User) (domain.User, error) {
	panic("unimplemented")
}

// DeleteUser implements domain.UserService.
func (u *userService) DeleteUser(id int) error {
	panic("unimplemented")
}

// GetUserByID implements domain.UserService.
func (u *userService) GetUserByID(id int) (domain.User, error) {
	panic("unimplemented")
}

// GetUsers implements domain.UserService.
func (u *userService) GetUsers() ([]domain.User, error) {
	panic("unimplemented")
}

// UpdateUser implements domain.UserService.
func (u *userService) UpdateUser(id int, user domain.User) (domain.User, error) {
	panic("unimplemented")
}
