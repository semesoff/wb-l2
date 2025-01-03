package main

import (
	"dev11/internal/initializer"
	"fmt"
	"log"
	"net/http"
)

func main() {
	init := initializer.NewInitializer()
	mux := init.Init()

	cfgApp := init.ConfigProvider.GetConfig().App
	address := fmt.Sprintf("%s:%s", cfgApp.Host, cfgApp.Port)
	log.Printf("Starting server on %s", address)

	// Start server
	if err := http.ListenAndServe(address, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
