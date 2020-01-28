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
	GetAllEventsOverlaping(events []domain.Event) []domain.EventsOverlaping
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

func (eService EventService) GetAllEventsOverlaping(events []domain.Event) ([]domain.EventsOverlaping, error) {
	var eventsOverlaping []domain.EventsOverlaping
	if _, err := eService.IsEventsValid(events); err != nil {
		return eventsOverlaping, errors.Wrap(err, "error when invoking OverlapingEvents")
	}

	return eService.calculateOverlapingEvents(events)
}

func (eService EventService) calculateOverlapingEvents(events []domain.Event) ([]domain.EventsOverlaping, error) {
	var eventsOverlaping []domain.EventsOverlaping
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
