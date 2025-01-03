package db

import (
	"dev11/internal/domain"
)

func (d *Database) AddEvent(event domain.Event) (domain.Event, error) {
	_, err := d.db.Exec("INSERT INTO events (title, date, user_id) VALUES ($1, $2, $3)", event.Title, event.Date, event.UserID)
	if err != nil {
		return domain.Event{}, err
	}
	return event, nil
}

func (d *Database) UpdateEvent(event domain.Event) (domain.Event, error) {
	_, err := d.db.Exec("UPDATE events SET title = $1, date = $2, user_id = $3 WHERE id = $4", event.Title, event.Date, event.UserID, event.ID)
	if err != nil {
		return domain.Event{}, err
	}
	return event, nil
}

func (d *Database) DeleteEvent(event domain.Event) error {
	_, err := d.db.Exec("DELETE FROM events WHERE id = $1", event.ID)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) GetEventsForDay(date string) ([]domain.Event, error) {
	rows, err := d.db.Query("SELECT id, title, date, user_id FROM events WHERE date = $1", date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []domain.Event
	for rows.Next() {
		var event domain.Event
		err := rows.Scan(&event.ID, &event.Title, &event.Date, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (d *Database) GetEventsForWeek(date string) ([]domain.Event, error) {
	rows, err := d.db.Query("SELECT id, title, date, user_id FROM events WHERE date >= $1 AND date < $1 + 7", date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []domain.Event
	for rows.Next() {
		var event domain.Event
		err := rows.Scan(&event.ID, &event.Title, &event.Date, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (d *Database) GetEventsForMonth(date string) ([]domain.Event, error) {
	rows, err := d.db.Query("SELECT id, title, date, user_id FROM events WHERE date >= $1 AND date < $1 + 30", date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []domain.Event
	for rows.Next() {
		var event domain.Event
		err := rows.Scan(&event.ID, &event.Title, &event.Date, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}
