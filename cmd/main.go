package main

import (
	"log"
	"net/http"
)

func main() {

	// load routers
	SetupRoutes()

	log.Printf("Starting on port %d\n", 8282)
	http.ListenAndServe(":8282", nil)
}
