package main

import (
	"log"
	"net/http"

	"github.com/darkphotonKN/go-ollama-chat/config"
)

func main() {

	// setup config
	cfg := config.LoadConfig()

	// load routers
	SetupRoutes(&cfg)

	log.Printf("Starting on port %d\n", 8282)
	http.ListenAndServe(":8282", nil)
}
