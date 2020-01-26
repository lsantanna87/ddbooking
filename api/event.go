package api

import (
	"fmt"

	"github.com/lsantanna87/ddbooking/pkg/domain"
)

type EventAPI struct {
}

type EventAPIInterface interface {
	GetOverlapingEvents(c domain.Calendar) []domain.EventsOverlaping
	AreEventsValid() []domain.Event
}

func (e EventAPI) GetOverlapingEvents(c domain.Calendar) []domain.EventsOverlaping {
	fmt.Println("Invoking")
	return []domain.EventsOverlaping{}
}

func (e EventAPI) AreEventsValid() []domain.Event {

	fmt.Println("Invoking")
	return []domain.Event{}
}
