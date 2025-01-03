package router

import (
	"dev11/internal/handler"
	"dev11/internal/middleware"
	"net/http"
)

type Router struct {
	mux          *http.ServeMux
	eventHandler handler.EventHandlerProvider
}

func NewRouter(mux *http.ServeMux, eventHandler handler.EventHandlerProvider) *Router {
	return &Router{
		mux:          mux,
		eventHandler: eventHandler,
	}
}

func (r *Router) Init() {
	mux := r.mux
	mux.Handle("/create_event", middleware.Logging(http.HandlerFunc(r.eventHandler.CreateEvent)))
	mux.Handle("/update_event", middleware.Logging(http.HandlerFunc(r.eventHandler.UpdateEvent)))
	mux.Handle("/delete_event", middleware.Logging(http.HandlerFunc(r.eventHandler.DeleteEvent)))
	mux.Handle("/events_for_day", middleware.Logging(http.HandlerFunc(r.eventHandler.GetEventsForDay)))
	mux.Handle("/events_for_week", middleware.Logging(http.HandlerFunc(r.eventHandler.GetEventsForWeek)))
	mux.Handle("/events_for_month", middleware.Logging(http.HandlerFunc(r.eventHandler.GetEventsForMonth)))
}
