package domain

import "github.com/savinnsk/api-template-go/internal/infra/sqlc"

type User struct {
	ID       string 
	Name     string 
	Email    string 
	Password string
	IsActive bool
	IsAdmin  bool
	Phone string
}

type UserDto struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Password string `json:"password" validate:"required,min=6"`
	Email    string `json:"email" validate:"required,email"`
	Phone string	`json:"phone,omitempty" validate:"min=10"`
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


func ToUserResponseList(userSqlcList *[]sqlc.User) *[]User{
	users := make([]User,0,len(*userSqlcList))

	for _,u := range *userSqlcList {
		users = append(users, User{
			ID : u.ID,
			Name : u.Name,
			Email: u.Email,
			Password: "",
		})
	}

	return &users
}

type LoginResponse struct {
    Token string `json:"token"`
}