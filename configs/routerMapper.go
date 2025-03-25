package configs

import (
	"net/http"
	userControler "github.com/savinnsk/api-template-go/internal/presentation"
)

func RouterMapper(mux *http.ServeMux) {
	mux.HandleFunc("GET /users/{id}",userControler.GetUserById)
}