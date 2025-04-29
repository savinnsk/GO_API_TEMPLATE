package domain

import "github.com/savinnsk/api-template-go/internal/infra/sqlc"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDto struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,email"`
}


type LoginDto struct {
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,email"`
}

func ToUserReponse(userSqlc *sqlc.User) *User{
	return &User{
		ID: userSqlc.ID,
		Name : userSqlc.Name,
		Email: userSqlc.Email,
		Password: userSqlc.Password,
	}
}