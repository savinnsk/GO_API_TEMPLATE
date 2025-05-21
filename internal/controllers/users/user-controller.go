package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/savinnsk/api-template-go/configs"
	"github.com/savinnsk/api-template-go/internal/controllers/middlewares"
	"github.com/savinnsk/api-template-go/internal/domain"
	"github.com/savinnsk/api-template-go/internal/shared"
	"github.com/savinnsk/api-template-go/internal/usecases/users"

)

//protected
func GetUserById(w http.ResponseWriter, r *http.Request) {
	
	id := r.PathValue("id")

	result,err := usecase.FindUserByIdUseCase(id)

	if err != nil {
		shared.BadRequest(w,err.Error())
	}

	 shared.Ok(w,map[string]interface{}{
		"message": "User Found",
		"user":    result,
	})
}

//protected 
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	

	result,err := usecase.GetAllUsersUseCase()

	if err != nil {
		shared.BadRequest(w,err.Error())
	}

	 shared.Ok(w,map[string]interface{}{
		"message": "Users Found",
		"user":    result,
	})
}

//protected 
func UpdateUserUser(w http.ResponseWriter, r *http.Request){

	claims, _ := r.Context().Value(middlewares.ClaimsKey).(jwt.MapClaims)
	
	subEmail := claims["sub"].(string)

	var user domain.UserDto
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		shared.BadRequest(w,err.Error())
		return
	}

	if err := configs.Validate.Struct(user); err != nil {
		shared.BadRequest(w, "Validation error: "+err.Error())
		return
	}

	result,err := usecase.UpdateUserUseCase(subEmail,user)

	if err != nil {
		if err.Error() == "internal error" {shared.InternalError(w,err.Error());return }
		shared.BadRequest(w,err.Error())
		return
	}
	
	
	 shared.Ok(w,map[string]interface{}{
		"message": "User updated",
		"user":    result,
	})


}

func CreateUser(w http.ResponseWriter, r *http.Request){
	
	var user domain.UserDto
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		shared.BadRequest(w,err.Error())
		return
	}

	if err := configs.Validate.Struct(user); err != nil {
		shared.BadRequest(w, "Validation error: "+err.Error())
		return
	}

	result,err := usecase.CreateUserUseCase(user)

	if err != nil {
		if err.Error() == "internal error" {shared.InternalError(w,err.Error());return }
		shared.BadRequest(w,err.Error())
		return
	}
	
	
	 shared.Ok(w,map[string]interface{}{
		"message": "User created successfully",
		"user":    result,
	})


}

func Login(w http.ResponseWriter, r *http.Request) {

	var user domain.LoginDto
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		shared.BadRequest(w,err.Error())
		return
	}

	result,err := usecase.LoginUserUseCase(user)

		if err != nil {
		if err.Error() == "internal error" {shared.InternalError(w,err.Error());return }
		shared.BadRequest(w,err.Error())
		return
	}

	 shared.Ok(w,map[string]interface{}{
		"message": "login successfully",
		"user":    result,
	})
	
}