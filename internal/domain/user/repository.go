package user

type Repository interface {
	Create(Entity *Entity) error
	GetByEmail(email string) (*Entity, error)
	GetByID(id uint) (*Entity, error)
	GetAllEntitys() ([]Entity, error)
}
