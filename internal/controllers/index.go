package controllers

import (
	"net/http"

	"github.com/savinnsk/api-template-go/internal/controllers/users"

)

func UserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /users/{id}",controllers.GetUserById)
	mux.HandleFunc("POST /users",controllers.CreateUser)
	mux.HandleFunc("POST /Login",controllers.Login)
}