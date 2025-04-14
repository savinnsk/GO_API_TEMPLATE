package configs

import (
	"net/http"

	"github.com/savinnsk/api-template-go/internal/presentation/users"

)

func RouterMapper(mux *http.ServeMux) {
	presentation.UserRoutes(mux)
}