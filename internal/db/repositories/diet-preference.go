package repositories

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DietPreference struct {
	Id         uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	Preference string    `gorm:"column:preference;not null" json:"preference"`
	UserId     int64     `gorm:"column:user_id;not null" json:"user_id"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

type DietPreferenceRepository struct {
	db *gorm.DB
}

func (r *DietPreferenceRepository) FindByUserId(userId int64) ([]DietPreference, error) {
	var preferences []DietPreference
	err := r.db.Where("user_id = ?", userId).Find(&preferences).Error
	return preferences, err
}

func (r *DietPreferenceRepository) Create(data []*DietPreference) error {
	return r.db.Create(&data).Error
}

func (r *DietPreferenceRepository) Save(preference *DietPreference) (*DietPreference, error) {
	err := r.db.Save(preference).Error
	return preference, err
}

func (r *DietPreferenceRepository) Remove(id int) error {
	err := r.db.Delete(&DietPreference{}, "id", id).Error
	return err
}

func (r *DietPreferenceRepository) RemoveByUserId(userId int64) error {
	err := r.db.Delete(&DietPreference{}, "user_id", userId).Error
	return err
}

func (p *DietPreference) TableName() string {
	return "diet_preferences"
}

func (p *DietPreference) BeforeCreate(tx *gorm.DB) (err error) {
	p.Id = uuid.New()
	return nil
}

func NewDietPreferenceRepository(db *gorm.DB) *DietPreferenceRepository {
	return &DietPreferenceRepository{db: db}
}
