package main

import (
	"log"
	"net/http"
	"github.com/savinnsk/api-template-go/configs"
)



func main() {
	cfg, err := configs.LoadConfig("../")
	if (err != nil) {panic(err)}

	db := configs.InitDb(cfg)
	defer db.Close()
	log.Printf("Database initialized with success")

	mux := http.NewServeMux()
	server := configs.LoadMiddlewares(mux)
	configs.RouterMapper(mux)

	log.Printf("Server Started At %s\n", cfg.PortServer)
	http.ListenAndServe(cfg.PortServer,server)
}
