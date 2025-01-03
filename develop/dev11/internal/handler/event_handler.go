package handler

import (
	"dev11/internal/domain"
	"dev11/internal/service"
	"dev11/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type EventHandlerProvider interface {
	CreateEvent(w http.ResponseWriter, r *http.Request)
	UpdateEvent(w http.ResponseWriter, r *http.Request)
	DeleteEvent(w http.ResponseWriter, r *http.Request)
	GetEventsForDay(w http.ResponseWriter, r *http.Request)
	GetEventsForWeek(w http.ResponseWriter, r *http.Request)
	GetEventsForMonth(w http.ResponseWriter, r *http.Request)
}

type EventHandler struct {
	service *service.EventServiceProvider
}

func NewEventHandler(service *service.EventServiceProvider) *EventHandler {
	return &EventHandler{
		service: service,
	}
}

func parseEvent(r *http.Request) (domain.Event, error) {
	var event domain.Event
	if r.Method == http.MethodPost {
		err := json.NewDecoder(r.Body).Decode(&event)
		fmt.Println(err, event)
		if err != nil {
			return domain.Event{}, err
		}
	} else {
		query := r.URL.Query()
		userID, err := strconv.Atoi(query.Get("user_id"))
		if err != nil {
			return domain.Event{}, err
		}
		date, err := time.Parse("2006-01-02", query.Get("date"))
		if err != nil {
			return domain.Event{}, err
		}
		id, _ := strconv.Atoi(query.Get("id"))
		event = domain.Event{
			ID:     id,
			UserID: userID,
			Date:   date,
			Title:  query.Get("title"),
		}
	}
	return event, nil
}

func (h *EventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": "invalid request method"})
		return
	}
	event, err := parseEvent(r)
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request data"})
		return
	}
	createdEvent, err := (*h.service).CreateEvent(event)
	if err != nil {
		utils.RespondJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]string{"result": fmt.Sprintf("%d", createdEvent.ID)})
}

func (h *EventHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": "invalid request method"})
		return
	}
	event, err := parseEvent(r)
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request data"})
		return
	}
	updatedEvent, err := (*h.service).UpdateEvent(event)
	if err != nil {
		utils.RespondJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]string{"result": fmt.Sprintf("%d", updatedEvent.ID)})
}

func (h *EventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": "invalid request method"})
		return
	}
	event, err := parseEvent(r)
	if err != nil {
		utils.RespondJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request data"})
		return
	}
	if err := (*h.service).DeleteEvent(event); err != nil {
		utils.RespondJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]string{"result": fmt.Sprintf("%d", event.ID)})
}

func (h *EventHandler) GetEventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": "invalid request method"})
		return
	}
	date := r.URL.Query().Get("date")
	events, err := (*h.service).GetEventsForDay(date)
	if err != nil {
		utils.RespondJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{"result": events})
}

func (h *EventHandler) GetEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": "invalid request method"})
		return
	}
	date := r.URL.Query().Get("date")
	events, err := (*h.service).GetEventsForWeek(date)
	if err != nil {
		utils.RespondJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{"result": events})
}

func (h *EventHandler) GetEventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.RespondJSON(w, http.StatusInternalServerError, map[string]string{"error": "invalid request method"})
		return
	}
	date := r.URL.Query().Get("date")
	events, err := (*h.service).GetEventsForMonth(date)
	if err != nil {
		utils.RespondJSON(w, http.StatusServiceUnavailable, map[string]string{"error": err.Error()})
		return
	}
	utils.RespondJSON(w, http.StatusOK, map[string]interface{}{"result": events})
}
