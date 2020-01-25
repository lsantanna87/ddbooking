package domain

import "time"

type Event struct {
	Name        string
	StartDate   time.Time
	EndDate     time.Time
	Description string
}

type EventsOverlaping struct {
	FirstEvent  Event
	SecondEvent Event
}
