package usecase

import "core/internal/domain"

type userUseCase struct {
	userService domain.UserService
}

func NewUserUseCase(userService domain.UserService) domain.UserUseCase {
	return &userUseCase{
		userService: userService,
	}
}

// CreateUser implements domain.UserUseCase.
func (u *userUseCase) CreateUser(user domain.User) (domain.User, error) {
	panic("unimplemented")
}

// DeleteUser implements domain.UserUseCase.
func (u *userUseCase) DeleteUser(id int) error {
	panic("unimplemented")
}

// GetUserByID implements domain.UserUseCase.
func (u *userUseCase) GetUserByID(id int) (domain.User, error) {
	panic("unimplemented")
}

// GetUsers implements domain.UserUseCase.
func (u *userUseCase) GetUsers() ([]domain.User, error) {
	panic("unimplemented")
}

// UpdateUser implements domain.UserUseCase.
func (u *userUseCase) UpdateUser(id int, user domain.User) (domain.User, error) {
	panic("unimplemented")
}
