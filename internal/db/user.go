package db

type User struct {
}

type UserRepository struct {
	db any
}

func NewUserRepository(db any) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindById(id string) (*User, error) {
	return nil, nil
}
