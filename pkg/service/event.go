package service

import (
	"sort"

	"github.com/lsantanna87/ddbooking/pkg/domain"
)

func SortEventByStartDate(events []domain.Event) []domain.Event {
	sortFunc := func(i, j int) bool {
		return events[i].StartDate.Before(events[j].StartDate)
	}
	sort.Slice(events, sortFunc)

	return events
}

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

func AllOverlapingEvents(cal domain.Calendar) []domain.EventsOverlaping {
	var eventsOverlaping []domain.EventsOverlaping

	cal.Events = SortEventByStartDate(cal.Events)

	for i := 0; i < len(cal.Events)-1; i++ {
		current := cal.Events[i]

		for j := i + 1; j < len(cal.Events); j++ {
			next := cal.Events[j]

			if currentOverlap, ok := GetEventIfOverlaping(current, next); ok {
				eventsOverlaping = append(eventsOverlaping, currentOverlap)
			}
		}
	}

	return eventsOverlaping
}
