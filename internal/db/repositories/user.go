package repositories

import "gorm.io/gorm"

type User struct {
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindById(id string) (*User, error) {
	return nil, nil
}
