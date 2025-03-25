package presentation

import (
	"fmt"
	"net/http"
)



func GetUserById(w http.ResponseWriter, r *http.Request) {
	id:= r.PathValue(("id"))

	fmt.Fprintf(w,"hello id:%s", id)
}