package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/savinnsk/api-template-go/internal/usecases/users"
	"github.com/savinnsk/api-template-go/internal/domain"

)



func GetUserById(w http.ResponseWriter, r *http.Request) {
	id:= r.PathValue(("id"))

	fmt.Fprintf(w,"hello id:%s", id)
}

func CreateUser(w http.ResponseWriter, r *http.Request){
	
	var user domain.UserDto
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result,err := usecase.CreateUserUseCase(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "User created successfully",
		"user":    result,
	})


}