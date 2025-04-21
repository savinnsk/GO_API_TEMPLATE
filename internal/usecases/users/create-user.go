package usecase

import (
	"github.com/savinnsk/api-template-go/internal/domain"
	"github.com/savinnsk/api-template-go/internal/infra/sqlc"
	"github.com/savinnsk/api-template-go/internal/infra/repositories"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

)

func CreateUserUseCase(userDto domain.UserDto )(*domain.User, error){

	hash,err := bcrypt.GenerateFromPassword([]byte(userDto.Password),bcrypt.DefaultCost)
	if err != nil {return nil,err}

	user := sqlc.CreateUserParams{
		ID: uuid.New().String(),
		Name: userDto.Name,
		Password:string(hash),
		Email: userDto.Email,
	}

	userSqlc,err := repositories.CreateUser(user)
	if err != nil {
		return nil,err
	}

	result := domain.ToUserReponse(&userSqlc)

	return result,nil

	
}