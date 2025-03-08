package dto

type CreateBotStateDto struct {
	ChatId int64
	State  string
}

type CreateDietPreferenceDto struct {
	UserId     int64  `json:"user_id"`
	Preference string `json:"preference"`
}
