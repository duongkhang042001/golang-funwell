package user

type Usecase interface {
	CreateUser(name, email, password string) error
	GetUserByEmail(email string) (*Entity, error)
	GetUserByID(id uint) (*Entity, error)
	GetAllUsers() ([]Entity, error)
	DeleteUser(id uint) error
	UpdateUser(id uint, name, email, password string) error
}
