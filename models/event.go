package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime" binding:"required"`
	UserID      int       `json:"user_id"`
}

var events = []Event{}

func (e Event) Save() {
	// save the event to the database
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
