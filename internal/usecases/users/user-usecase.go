package usecase

import (
	"fmt"
	"log"

	"github.com/savinnsk/api-template-go/internal/domain"
	infra "github.com/savinnsk/api-template-go/internal/infra/jwt"
	"github.com/savinnsk/api-template-go/internal/infra/repositories"
	"golang.org/x/crypto/bcrypt"

)

func CreateUserUseCase(userDto domain.UserDto )(*domain.User, error){

	hash,err := bcrypt.GenerateFromPassword([]byte(userDto.Password),bcrypt.DefaultCost)
	if err != nil {
		log.Printf("CreateUserUseCase.GenerateFromPassword: %v", err)
		return nil,fmt.Errorf("internal error",)
	}

	userDto.Password = string(hash)

	userSqlc,err := repositories.CreateUser(userDto)
	if err != nil {
		log.Printf("CreateUserUseCase.CreateUser: %v", err)
		return nil,fmt.Errorf("internal error",)
	}

	result := domain.ToUserReponse(&userSqlc)

	return result,nil

	
}

func LoginUserUseCase(loginDto domain.LoginDto)(*domain.LoginResponse, error){

	user ,err := repositories.GetByEmail(loginDto.Email)
	if err != nil {
		log.Printf("LoginUserUseCase.GetByEmail: %v", err)
		return nil,fmt.Errorf("email or password invalid")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(loginDto.Password))
		if err != nil {
		log.Printf("LoginUserUseCase.CompareHashAndPassword: %v", err)	
		return nil,fmt.Errorf("email or password invalid")
	}

	token ,err := infra.GenerateJwt(user.Email)
	if err != nil {
		log.Printf("LoginUserUseCase.GenerateJwt: %v", err)
		return nil,fmt.Errorf("internal error")
	}


	return token,nil

}

func FindUserByIdUseCase(id string)(*domain.User, error){

	userSqlc ,err := repositories.GetUserById(id)
	if err != nil {
		log.Printf("FindUserByIdUseCase.GetUserById error: %v", err)
		return nil,fmt.Errorf("user not found")
	}

	result := domain.ToUserReponse(&userSqlc)

	result.Password = ""


	return result,nil

}

func GetAllUsersUseCase()(*[]domain.User, error){

	usersSqlc ,err := repositories.GetUsers()
	if err != nil {
		log.Printf("GetAllUsersUseCase.GetUsers error: %v", err)
		return nil,fmt.Errorf("internal error")
	}

	result := domain.ToUserResponseList(&usersSqlc)

	return result,nil

}

func UpdateUserUseCase(subEmail string,userDto domain.UserDto)(*domain.User, error){

	userSqlc ,err := repositories.GetByEmail(subEmail)
		if err != nil || userSqlc.Email != subEmail{
		log.Printf("UpdateUserUseCase.GetByEmail error: %v", err)
		return nil,fmt.Errorf("internal error")
	}

	hash,err := bcrypt.GenerateFromPassword([]byte(userDto.Password),bcrypt.DefaultCost)
	if err != nil {
		log.Printf("CreateUserUseCase.GenerateFromPassword: %v", err)
		return nil,fmt.Errorf("internal error",)
	}

	userSqlc.Name = userDto.Name
	userSqlc.Email= userDto.Email
	userSqlc.Password = string(hash)

	userSqlcUpdated ,err := repositories.UpdateUser(userSqlc)
	if err != nil {
		log.Printf("UpdateUser.UpdateUser error: %v", err)
		return nil,fmt.Errorf("internal error")
	}

	result := domain.ToUserReponse(&userSqlcUpdated)

	result.Password = ""


	return result,nil

}