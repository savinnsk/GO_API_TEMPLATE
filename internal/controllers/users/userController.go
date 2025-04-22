package controllers

import (
	"fmt"
	"net/http"

)



func GetUserById(w http.ResponseWriter, r *http.Request) {
	id:= r.PathValue(("id"))

	fmt.Fprintf(w,"hello id:%s", id)
}

func CreateUser(w http.ResponseWriter, r *http.Request){

	//body , _ := r.GetBody();


}