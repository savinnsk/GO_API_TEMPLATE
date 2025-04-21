package domain

import "github.com/savinnsk/api-template-go/internal/infra/sqlc"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDto struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToUserReponse(userSqlc *sqlc.User) *User{
	return &User{
		ID: userSqlc.ID,
		Name : userSqlc.Name,
		Email: userSqlc.Email,
		Password: userSqlc.Password,
	}
}