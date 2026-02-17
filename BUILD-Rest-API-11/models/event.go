package mod

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	userID      int
}

var events = []Event{}

func (e Event) save() {
	// laterÑ ad it to a database
	events = append(events)
}
