package repositories

import (
	"core/internal/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(dbConn *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: dbConn,
	}
}

// Create implements domain.UserRepository.
func (u *userRepository) Create(user domain.User) (domain.User, error) {
	panic("unimplemented")
}

// GetAll implements domain.UserRepository.
func (u *userRepository) GetAll() ([]domain.User, error) {
	panic("unimplemented")
}

// GetByID implements domain.UserRepository.
func (u *userRepository) GetByID(id int) (domain.User, error) {
	panic("unimplemented")
}

// Update implements domain.UserRepository.
func (u *userRepository) Update(user domain.User) (domain.User, error) {
	panic("unimplemented")
}

// Delete implements domain.UserRepository.
func (u *userRepository) Delete(id int) error {
	panic("unimplemented")
}
