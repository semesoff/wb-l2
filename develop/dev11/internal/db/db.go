package db

import (
	"database/sql"
	"dev11/config"
	"dev11/internal/domain"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type DatabaseProvider interface {
	AddEvent(event domain.Event) (domain.Event, error)
	UpdateEvent(event domain.Event) (domain.Event, error)
	DeleteEvent(event domain.Event) error
	GetEventsForDay(date string) ([]domain.Event, error)
	GetEventsForWeek(date string) ([]domain.Event, error)
	GetEventsForMonth(date string) ([]domain.Event, error)
}

type Database struct {
	db *sql.DB
}

func NewDatabase(cfg config.Database) *Database {
	db := &Database{}
	db.Init(cfg)
	return db
}

func (d *Database) Init(cfg config.Database) {
	db, err := sql.Open(cfg.Driver,
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name))
	if err != nil {
		log.Fatalf("Failed to open a DB connection: %v", err)
	}
	d.db = db
	log.Println("Database is initialized.")
}
