package dto

type CreateBotStateDto struct {
	ChatId int64
	State  string
}

type CreateDietPreferenceDto struct {
	UserId     int64  `json:"user_id"`
	Preference string `json:"preference"`
}

type CreateSavedRecipeDto struct {
	MessageId int    `json:"message_id"`
	Photo     string `json:"photo"`
	Caption   string `json:"caption"`
	UserId    int64  `json:"user_id"`
}
