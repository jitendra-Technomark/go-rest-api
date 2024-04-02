package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string     `binding:"required"`
	Description string     `binding:"required"`
	Location    string     `binding:"required"`
	DateTime    *time.Time `binding:"required"`
	UserID      int
}

func (e Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, dateTime, user_id) 
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
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		var dateTime []uint8 // Temporary variable to hold dateTime value
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &dateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		// Convert []uint8 to time.Time
		parsedTime, err := time.Parse("2006-01-02 15:04:05", string(dateTime))
		if err != nil {
			return nil, err
		}

		// Assign parsed time to the DateTime pointer
		event.DateTime = &parsedTime

		events = append(events, event)
	}

	return events, nil
}
