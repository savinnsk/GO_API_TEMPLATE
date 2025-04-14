package presentation

import (
	"net/http"

)

func UserRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /users/{id}",GetUserById)
}