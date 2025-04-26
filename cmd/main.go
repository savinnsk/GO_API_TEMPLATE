package main

import (
	"log"
	"net/http"

	"github.com/savinnsk/api-template-go/configs"
	controllers "github.com/savinnsk/api-template-go/internal/controllers"

)

func main() {
	env, err := configs.LoadEnv("../")
	if err != nil {
		panic(err)
	}

	db := configs.InitDb(env)
	defer db.Close()
	log.Printf("Database initialized with success")

	configs.LoadContext()

	mux := http.NewServeMux()
	server := configs.LoadMiddlewares(mux)
	controllers.RouterMapper(mux)

	log.Printf("Server Started At %s\n", env.PortServer)
	http.ListenAndServe(env.PortServer, server)
}
