package repositories

import (
	"time"

	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/db/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SavedRecipe struct {
	Id        uuid.UUID `gorm:"column:id;primaryKey" json:"id"`
	MessageId int       `gorm:"column:message_id;unique;not null" json:"message_id"`
	Photo     string    `gorm:"column:photo;not null" json:"photo"`
	Caption   string    `gorm:"column:caption;not null" json:"caption"`
	UserId    int64     `gorm:"column:user_id;not null" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

type SavedRecipeRepository struct {
	db *gorm.DB
}

func (s *SavedRecipeRepository) FindByUserId(userId int64) ([]SavedRecipe, error) {
	var recipes []SavedRecipe
	err := s.db.Where("user_id = ?", userId).Find(&recipes).Error
	return recipes, err
}

func (s *SavedRecipeRepository) Create(data []*SavedRecipe) error {
	return s.db.Create(&data).Error
}

func (s *SavedRecipeRepository) Upsert(data dto.CreateSavedRecipeDto) error {
	return s.db.Clauses(clause.OnConflict{
		Columns: []clause.Column{
			{Name: "message_id"},
		},
		DoUpdates: clause.AssignmentColumns([]string{"photo", "caption", "user_id"}),
	}).Create(&SavedRecipe{
		MessageId: data.MessageId,
		Photo:     data.Photo,
		Caption:   data.Caption,
		UserId:    data.UserId,
	}).Error
}

func (s *SavedRecipeRepository) Save(recipe *SavedRecipe) (*SavedRecipe, error) {
	err := s.db.Save(recipe).Error
	return recipe, err
}

func (s *SavedRecipeRepository) Remove(id string) error {
	err := s.db.Unscoped().Delete(&SavedRecipe{}, "id", id).Error
	return err
}

func (s *SavedRecipeRepository) RemoveByUserId(userId int64) error {
	err := s.db.Unscoped().Delete(&SavedRecipe{}, "user_id", userId).Error
	return err
}

func (s *SavedRecipeRepository) RemoveByPhoto(photo string) error {
	err := s.db.Unscoped().Delete(&SavedRecipe{}, "photo", photo).Error
	return err
}

func NewSavedRecipeRepository(db *gorm.DB) *SavedRecipeRepository {
	return &SavedRecipeRepository{db: db}
}

func (s *SavedRecipe) BeforeCreate(tx *gorm.DB) (err error) {
	s.Id = uuid.New()
	return
}
