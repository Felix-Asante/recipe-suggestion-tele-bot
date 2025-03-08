package repositories

import (
	"errors"

	"github.com/Felix-Asante/recipe-suggestion-tele-bot/internal/db/dto"
	"gorm.io/gorm"
)

type BotState struct {
	gorm.Model
	ChatID int64  `gorm:"column:chat_id;unique" json:"chat_id"`
	State  string `gorm:"column:state" json:"state"`
}

type BotStateRepository struct {
	db *gorm.DB
}

func (s *BotStateRepository) FindByChatId(chatId int64) (*BotState, error) {
	var state BotState
	err := s.db.First(&state, "chat_id", chatId).Error
	return &state, err
}

func (s *BotStateRepository) Create(data dto.CreateBotStateDto) (*BotState, error) {
	state := &BotState{
		ChatID: data.ChatId,
		State:  data.State,
	}
	return s.Save(state)
}

func (s *BotStateRepository) Upsert(data dto.CreateBotStateDto) (*BotState, error) {
	state, err := s.FindByChatId(data.ChatId)
	if nil != err && !errors.Is(err, gorm.ErrRecordNotFound) {
		return s.Create(data)
	}
	state.State = data.State
	state.ChatID = data.ChatId
	return s.Save(state)
}

func (s *BotStateRepository) Save(state *BotState) (*BotState, error) {
	err := s.db.Save(state).Error
	return state, err
}

func (s *BotStateRepository) RemoveByChatId(chatId int64) error {
	err := s.db.Delete(&BotState{}, "chat_id", chatId).Error
	return err
}

func (s *BotStateRepository) Remove(id int) error {
	err := s.db.Delete(&BotState{}, "id", id).Error
	return err
}

func NewBotStateRepository(db *gorm.DB) *BotStateRepository {
	return &BotStateRepository{db: db}
}
func (s *BotState) TableName() string {
	return "bot_state"
}
