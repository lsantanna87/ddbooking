package service

import (
	"sort"

	"github.com/lsantanna87/ddbooking/pkg/domain"
)

func GetEventIfOverlaping(currentEvent domain.Event, nextEvent domain.Event) (domain.EventsOverlaping, bool) {
	if nextEvent.StartDate.Sub(currentEvent.EndDate) <= 0 {
		currentOverlap := domain.EventsOverlaping{
			FirstEvent:  currentEvent,
			SecondEvent: nextEvent,
		}

		return currentOverlap, true
	}

	return domain.EventsOverlaping{}, false
}

func AllOverlapingEvents(events []domain.Event) []domain.EventsOverlaping {
	var eventsOverlaping []domain.EventsOverlaping
	events = sortEventByStartDate(events)

	for i := 0; i < len(events)-1; i++ {
		current := events[i]

		for j := i + 1; j < len(events); j++ {
			next := events[j]
			if currentOverlap, ok := GetEventIfOverlaping(current, next); ok {
				eventsOverlaping = append(eventsOverlaping, currentOverlap)
			}
		}
	}

	return eventsOverlaping
}

func sortEventByStartDate(events []domain.Event) []domain.Event {
	sortFunc := func(i, j int) bool {
		return events[i].StartDate.Before(events[j].StartDate)
	}
	sort.Slice(events, sortFunc)

	return events
}
