package domain

type User struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
}

type UserUseCase interface {
	GetUsers() ([]User, error)
	GetUserByID(id int) (User, error)
	CreateUser(user User) (User, error)
	UpdateUser(id int, user User) (User, error)
	DeleteUser(id int) error
}

type UserService interface {
	GetUsers() ([]User, error)
	GetUserByID(id int) (User, error)
	CreateUser(user User) (User, error)
	UpdateUser(id int, user User) (User, error)
	DeleteUser(id int) error
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetByID(id int) (User, error)
	Create(user User) (User, error)
	Update(user User) (User, error)
	Delete(id int) error
}
