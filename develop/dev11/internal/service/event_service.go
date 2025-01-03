package service

import (
	"dev11/internal/db"
	"dev11/internal/domain"
	"errors"
)

type EventServiceProvider interface {
	CreateEvent(event domain.Event) (domain.Event, error)
	UpdateEvent(event domain.Event) (domain.Event, error)
	DeleteEvent(event domain.Event) error
	GetEventsForDay(date string) ([]domain.Event, error)
	GetEventsForWeek(date string) ([]domain.Event, error)
	GetEventsForMonth(date string) ([]domain.Event, error)
}

type EventService struct {
	db *db.DatabaseProvider
}

func NewEventService(db *db.DatabaseProvider) *EventService {
	return &EventService{db: db}
}

func (s *EventService) CreateEvent(event domain.Event) (domain.Event, error) {
	if event.UserID == 0 || event.Date.IsZero() || event.Title == "" {
		return event, errors.New("invalid event data")
	}
	createdEvent, err := (*s.db).AddEvent(event)
	return createdEvent, err
}

func (s *EventService) UpdateEvent(event domain.Event) (domain.Event, error) {
	if event.Date.IsZero() || event.Title == "" {
		return event, errors.New("invalid event data")
	}
	updatedEvent, err := (*s.db).UpdateEvent(event)
	return updatedEvent, err
}

func (s *EventService) DeleteEvent(event domain.Event) error {
	if event.ID < 0 {
		return errors.New("invalid event data")
	}
	err := (*s.db).DeleteEvent(event)
	return err
}

func (s *EventService) GetEventsForDay(date string) ([]domain.Event, error) {
	if date == "" {
		return nil, errors.New("invalid date")
	}
	events, err := (*s.db).GetEventsForDay(date)
	return events, err
}

func (s *EventService) GetEventsForWeek(date string) ([]domain.Event, error) {
	if date == "" {
		return nil, errors.New("invalid date")
	}
	events, err := (*s.db).GetEventsForWeek(date)
	return events, err
}

func (s *EventService) GetEventsForMonth(date string) ([]domain.Event, error) {
	if date == "" {
		return nil, errors.New("invalid date")
	}
	events, err := (*s.db).GetEventsForMonth(date)
	return events, err
}
