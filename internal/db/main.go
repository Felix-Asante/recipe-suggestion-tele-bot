package db

type Storage struct {
	User interface {
		FindById(id string) (*User, error)
	}
}

func NewStorage() *Storage {
	return &Storage{User: NewUserRepository(nil)}
}
