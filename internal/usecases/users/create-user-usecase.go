package usecase

import (
	"github.com/savinnsk/api-template-go/internal/domain"
	"github.com/savinnsk/api-template-go/internal/infra/repositories"
	"golang.org/x/crypto/bcrypt"

)

func CreateUserUseCase(userDto domain.UserDto )(*domain.User, error){

	hash,err := bcrypt.GenerateFromPassword([]byte(userDto.Password),bcrypt.DefaultCost)
	if err != nil {return nil,err}

	userDto.Password = string(hash)

	userSqlc,err := repositories.CreateUser(userDto)
	if err != nil {
		return nil,err
	}

	result := domain.ToUserReponse(&userSqlc)

	return result,nil

	
}