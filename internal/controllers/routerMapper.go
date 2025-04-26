package controllers

import (
	"net/http"

)

func RouterMapper(mux *http.ServeMux) {
	UserRoutes(mux)
}