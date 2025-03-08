package dto

type CreateUserDto struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserId    int64  `json:"user_id"`
}
