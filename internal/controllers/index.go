package controllers

import (
	"net/http"

	"github.com/savinnsk/api-template-go/internal/controllers/middlewares"
	"github.com/savinnsk/api-template-go/internal/controllers/users"

)

func UserRoutes(mux *http.ServeMux) {
	mux.Handle("GET /users/{id}", middlewares.JwtAuth(http.HandlerFunc(controllers.GetUserById)))
	mux.HandleFunc("POST /users",controllers.CreateUser)
	mux.HandleFunc("POST /login",controllers.Login)
}