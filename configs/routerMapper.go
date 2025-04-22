package configs

import (
	"net/http"

	"github.com/savinnsk/api-template-go/internal/controllers/users"

)

func RouterMapper(mux *http.ServeMux) {
	controllers.UserRoutes(mux)
}