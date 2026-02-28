package models

import (
	"time"

	"example.com/BUILD-Rest-API-11/db"
)

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	// laterÑ ad it to a database
	query := `INSERT INTO events (name, description, location, date_time, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = int(id)

	return err
}

func GetAllEvents() []Event {
	query := `SELECT id, name, description, location, date_time, user_id FROM events`
	db.DB.Query(query)
	return events
}
