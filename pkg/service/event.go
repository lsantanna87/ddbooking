package service

import (
	"fmt"
	"sort"

	"github.com/lsantanna87/ddbooking/pkg/domain"
)

type EventService struct{}

type EventServiceInterface interface {
	OverlapingEvents(events []domain.Event) []domain.EventsOverlaping
	IsEventsValid(events []domain.Event) (bool, error)
}

func (eService EventService) OverlapingEvents(events []domain.Event) ([]domain.EventsOverlaping, error) {
	var eventsOverlaping []domain.EventsOverlaping

	if ok, err := eService.IsEventsValid(events); err != nil && !ok {
		return eventsOverlaping, err
	}

	events = eService.sortEventByStartDate(events)

	for i := 0; i < len(events)-1; i++ {
		current := events[i]
		for j := i + 1; j < len(events); j++ {
			next := events[j]
			if eService.isEventsOverlaping(current, next) {
				currentOverlap := domain.EventsOverlaping{FirstEvent: current, SecondEvent: next}
				eventsOverlaping = append(eventsOverlaping, currentOverlap)
			}
		}
	}

	return eventsOverlaping, nil
}

func (eService EventService) IsEventsValid(events []domain.Event) (bool, error) {
	var valid bool

	if len(events) <= 1 {
		return valid, fmt.Errorf("number of events has to be greater than 1.")
	}

	for _, v := range events {
		if ok, err := v.IsValid(); err != nil && !ok {
			return valid, err
		}
	}

	return valid, nil
}

func (eService EventService) isEventsOverlaping(currentEvent domain.Event, nextEvent domain.Event) bool {
	return nextEvent.StartDate.Sub(currentEvent.EndDate) <= 0
}

func (eService EventService) sortEventByStartDate(events []domain.Event) []domain.Event {
	sortFunc := func(i, j int) bool {
		return events[i].StartDate.Before(events[j].StartDate)

	}

	sort.Slice(events, sortFunc)

	return events
}
