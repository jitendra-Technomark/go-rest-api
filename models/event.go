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
	UserID      int64
}

// type Event struct {
// 	ID          int64      `db:"id" binding:"required"`
// 	Name        string     `db:"name" binding:"required"`
// 	Description string     `db:"description" binding:"required"`
// 	Location    string     `db:"location" binding:"required"`
// 	DateTime    *time.Time `db:"dateTime" binding:"required"`
// 	UserID      int        `db:"user_id"`
// }

func (e *Event) Save() error {
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
		var dateTime []uint8
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

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	var dateTime []uint8 // Temporary variable to hold dateTime value
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &dateTime, &event.UserID)

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

	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

// with sqlx package

// func (event Event) Update() error {
// 	query := `
// 	UPDATE events
// 	SET name = ?, description = ?, location = ?, dateTime = ?
// 	WHERE id = ?
// 	`

// 	_, err := db.DB.Exec(query, event.Name, event.Description, event.Location, event.DateTime, event.ID)
// 	return err
// }

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID)
	return err
}

// func (event Event) Delete() error {
// 	query := "DELETE FROM events WHERE id = :id"

// 	_, err := db.DB.NamedExec(query, map[string]interface{}{"id": event.ID})
// 	return err
// }
