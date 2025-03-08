package repositories

import (
	"errors"
	"time"

	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/db/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID  `gorm:"column:id;primaryKey" json:"id"`
	FirstName string     `gorm:"column:first_name;not null" json:"first_name"`
	LastName  string     `gorm:"column:last_name;not null" json:"last_name"`
	UserId    int64      `gorm:"column:user_id;unique;not null" json:"user_id"`
	CreatedAt time.Time  `gorm:"column:created_at;not null" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at;not null" json:"updated_at"`
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindById(id string) (*User, error) {
	var user User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) FindByUserId(userId int64) (*User, error) {
	var user User
	err := r.db.First(&user, "user_id", userId).Error
	return &user, err
}

func (r *UserRepository) Upsert(data dto.CreateUserDto) error {
	user, error := r.FindByUserId(data.UserId)

	if errors.Is(error, gorm.ErrRecordNotFound) {
		if _, err := r.Create(data); err != nil {
			return err
		}
		return nil
	}

	user.FirstName = data.FirstName
	user.LastName = data.LastName

	_, err := r.Save(user)
	return err
}

func (r *UserRepository) Create(data dto.CreateUserDto) (*User, error) {
	user := &User{
		ID:        uuid.New(),
		FirstName: data.FirstName,
		LastName:  data.LastName,
		UserId:    data.UserId,
	}

	return r.Save(user)
}

func (r *UserRepository) Save(user *User) (*User, error) {

	error := r.db.Save(user).Error

	return user, error
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	u.UpdatedAt = &time.Time{}
	return
}
