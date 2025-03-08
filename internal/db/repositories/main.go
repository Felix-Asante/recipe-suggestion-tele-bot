package repositories

import (
	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/db/dto"
	"gorm.io/gorm"
)

type Repositories struct {
	User interface {
		FindById(id string) (*User, error)
		FindByUserId(userId int64) (*User, error)
		Upsert(data dto.CreateUserDto) error
		Create(data dto.CreateUserDto) (*User, error)
		Save(user *User) (*User, error)
	}
}

func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{User: NewUserRepository(db)}
}
