package domain

import (
	"context"
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:255;not null"`
	Email     string `gorm:"size:255;unique;not null"`
	Password  string `gorm:"size:255;not null"`
	Role      string `gorm:"size:50;not null;default:'user'"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserPagination struct {
	Total    int
	PerPage  int
	Page     int
	LastPage int
	Data     []User
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindById(ctx context.Context, id uint) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id uint) error
	Paginate(ctx context.Context, page int, perPage int) (*UserPagination, error)
}

type UserService interface {
	Register(ctx context.Context, user *User) error
	Login(ctx context.Context, email, password string) (*User, error)
	FindById(ctx context.Context, id uint) (*User, error)
	Update(ctx context.Context, user *User) error
}

type UserUseCase interface {
}
