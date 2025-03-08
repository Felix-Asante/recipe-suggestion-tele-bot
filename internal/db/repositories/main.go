package repositories

import "gorm.io/gorm"

type Repositories struct {
	User interface {
		FindById(id string) (*User, error)
	}
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{User: NewUserRepository(db)}
}
