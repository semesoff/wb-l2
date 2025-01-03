package initializer

import (
	"dev11/config"
	"dev11/internal/db"
	"dev11/internal/handler"
	"dev11/internal/router"
	"dev11/internal/service"
	"net/http"
)

type Initializer struct {
	ConfigProvider       config.ConfigProvider
	DatabaseProvider     db.DatabaseProvider
	EventServiceProvider service.EventServiceProvider
	EventHandlerProvider handler.EventHandlerProvider
}

func NewInitializer() *Initializer {
	return &Initializer{}
}

func (i *Initializer) Init() *http.ServeMux {
	// Init http server
	mux := http.NewServeMux()

	// Init config
	i.ConfigProvider = config.NewConfigManager()
	cfg := i.ConfigProvider.GetConfig()

	// Init database
	i.DatabaseProvider = db.NewDatabase(cfg.Database)

	// Init event service
	i.EventServiceProvider = service.NewEventService(&i.DatabaseProvider)

	// Init event handler
	i.EventHandlerProvider = handler.NewEventHandler(&i.EventServiceProvider)

	// Init routers
	router.NewRouter(mux, i.EventHandlerProvider).Init()

	return mux
}
