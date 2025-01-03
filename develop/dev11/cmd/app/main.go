package main

import (
	"dev11/config"
	"dev11/internal/db"
	"dev11/internal/handler"
	"dev11/internal/router"
	"dev11/internal/service"
	"log"
	"net/http"
)

func main() {
	// Init http server
	mux := http.NewServeMux()

	// Init config
	var cfgProvider config.ConfigProvider = config.NewConfigManager()
	cfg := cfgProvider.GetConfig()

	// Init database
	var dbProvider db.DatabaseProvider = db.NewDatabase(cfg.Database)

	// Init event service
	var eventServiceProvider service.EventServiceProvider = service.NewEventService(&dbProvider)

	// Init event handler
	var eventHandlerProvider handler.EventHandlerProvider = handler.NewEventHandler(&eventServiceProvider)

	// Init routers
	router.NewRouter(mux, eventHandlerProvider).Init()

	address := cfg.App.Host + ":" + cfg.App.Port
	log.Println("Server is running on " + address)
	// Start server
	if err := http.ListenAndServe(address, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
