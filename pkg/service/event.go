package service

import (
	"fmt"
	"sort"

	"github.com/lsantanna87/ddbooking/pkg/domain"
	"github.com/pkg/errors"
)

type EventService struct{}

type EventServiceInterface interface {
	IsEventsValid(events []domain.Event) (bool, error)
	AllEventsOverlapping(events []domain.Event) []domain.EventsOverlapping
}

func (eService EventService) IsEventsValid(events []domain.Event) (bool, error) {
	if len(events) <= 1 {
		return false, fmt.Errorf("number of events has to be greater than 1.")
	}

	for _, v := range events {
		if _, err := v.IsValid(); err != nil {
			return false, err
		}
	}

	return true, nil
}

func (eService EventService) AllEventsOverlapping(events []domain.Event) ([]domain.EventsOverlapping, error) {
	if _, err := eService.IsEventsValid(events); err != nil {
		return []domain.EventsOverlapping{}, errors.Wrap(err, "error when invoking AllEventsOverlapping, events are not valid!")
	}

	return eService.calculateOverlappingEvents(events), nil
}

func (eService EventService) calculateOverlappingEvents(events []domain.Event) []domain.EventsOverlapping {
	var eventsOverlapping []domain.EventsOverlapping

	events = eService.sortEventByStartDate(events)

	for i := 0; i < len(events)-1; i++ {
		current := events[i]
		for j := i + 1; j < len(events); j++ {
			next := events[j]
			if eService.isEventsOverlapping(current, next) {
				eventsOverlapping = append(eventsOverlapping, domain.EventsOverlapping{FirstEvent: current, SecondEvent: next})
			}
		}
	}

	return eventsOverlapping
}

func (eService EventService) isEventsOverlapping(currentEvent domain.Event, nextEvent domain.Event) bool {
	return nextEvent.StartDate.Sub(currentEvent.EndDate) <= 0
}

func (eService EventService) sortEventByStartDate(events []domain.Event) []domain.Event {
	sortFunc := func(i, j int) bool {
		return events[i].StartDate.Before(events[j].StartDate)
	}

	sort.Slice(events, sortFunc)

	return events
}
