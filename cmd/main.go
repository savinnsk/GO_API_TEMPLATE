package main

import (
	"log"
	"net/http"
	"github.com/savinnsk/api-template-go/configs"
)



func main() {
	cfg, err := configs.LoadConfig("../")
	if (err != nil) {panic(err)}

	mux := http.NewServeMux()
	handler := configs.EnableCORS(mux)
	configs.RouterMapper(mux)

	log.Printf("Server Started At %s\n", cfg.PortServer)
	http.ListenAndServe(cfg.PortServer,handler)
}
