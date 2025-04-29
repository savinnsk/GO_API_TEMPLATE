package usecase

import (
	"fmt"

	"github.com/savinnsk/api-template-go/internal/domain"
	"github.com/savinnsk/api-template-go/internal/infra/repositories"
	"golang.org/x/crypto/bcrypt"

)

func CreateUserUseCase(userDto domain.UserDto )(*domain.User, error){

	hash,err := bcrypt.GenerateFromPassword([]byte(userDto.Password),bcrypt.DefaultCost)
	if err != nil {return nil,fmt.Errorf("internal error",)}

	userDto.Password = string(hash)

	userSqlc,err := repositories.CreateUser(userDto)
	if err != nil {
		return nil,fmt.Errorf("internal error",)
	}

	result := domain.ToUserReponse(&userSqlc)

	return result,nil

	
}

func LoginUser(loginDto domain.LoginDto)(*domain.User, error){

	user ,err := repositories.GetByEmail(loginDto.Email)
	if err != nil {
		return nil,fmt.Errorf("email or password invalid")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(loginDto.Password))
		if err != nil {
		return nil,fmt.Errorf("email or password invalid")
	}

	result := domain.ToUserReponse(&user)

	return result,nil

}