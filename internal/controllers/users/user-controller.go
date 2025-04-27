package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/savinnsk/api-template-go/configs"
	"github.com/savinnsk/api-template-go/internal/domain"
	"github.com/savinnsk/api-template-go/internal/shared"
	"github.com/savinnsk/api-template-go/internal/usecases/users"

)



func GetUserById(w http.ResponseWriter, r *http.Request) {
	id:= r.PathValue(("id"))

	fmt.Fprintf(w,"hello id:%s", id)
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
		shared.BadRequest(w,err.Error())
		return
	}
	
	 shared.Created(w,map[string]interface{}{
		"message": "User created successfully",
		"user":    result,
	})


}