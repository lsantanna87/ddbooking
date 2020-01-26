package api

import (
	"github.com/lsantanna87/ddbooking/pkg/domain"
	"github.com/lsantanna87/ddbooking/pkg/service"
)

type EventAPI struct {
	Events []domain.Event
}

type EventAPIInterface interface {
	GetOverlapingEvents(c domain.Event) []domain.EventsOverlaping
	IsEventValid() domain.Event
}

func (eAPI *EventAPI) GetOverlapingEvents() []domain.EventsOverlaping {
	return service.AllOverlapingEvents(eAPI.Events)
}

func (eAPI *EventAPI) IsEventValid() (isValid bool) {
	// for _, event := range c.Event.Events {
	// isValid = event.IsValid()
	// }
	return
}
