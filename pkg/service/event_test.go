package service

import (
	"strings"
	"testing"
	"time"

	"github.com/lsantanna87/ddbooking/pkg/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EventServiceTestSuite struct {
	suite.Suite
	t *testing.T
}

func (e *EventServiceTestSuite) SetupSuite() {
	e.t = e.T()
}
func TestEventServiceSuite(t *testing.T) {
	suite.Run(t, new(EventServiceTestSuite))
}

func (e *EventServiceTestSuite) TestShouldReturnTrueWhenEventsOverlapping() {
	event1 := domain.CreateFakeEvent()
	event2 := domain.CreateFakeEvent()
	event2.StartDate = event1.StartDate.Add(time.Minute * -10)
	event2.EndDate = event1.EndDate

	isOverlapping := EventService{}.isEventsOverlapping(event1, event2)

	assert.True(e.t, isOverlapping)
}

func (e *EventServiceTestSuite) TestShouldReturnFalseWhenEventsNotOverlapping() {
	event1 := domain.CreateFakeEvent()
	event2 := domain.CreateFakeEvent()
	event2.StartDate = event1.StartDate.Add(time.Minute * 60)
	event2.EndDate = event1.EndDate.Add(time.Minute * 120)

	isOverlapping := EventService{}.isEventsOverlapping(event1, event2)

	assert.False(e.t, isOverlapping)
}

func (e *EventServiceTestSuite) TestShouldReturnSortedEventsWhenSortEventByStartDate() {
	events := domain.CreateFakeEvents(50)
	events = EventService{}.sortEventByStartDate(events)

	for i := 0; i < len(events)-1; i++ {
		currentEvent := events[i]
		next := events[i+1]
		assert.True(e.t, currentEvent.StartDate.Before(next.StartDate))
	}
}

func (e *EventServiceTestSuite) TestShouldReturnErrorWhenCheckForOverlappingEventsWithSingleEvent() {
	events := domain.CreateFakeEvents(1)

	overlapping, err := EventService{}.AllEventsOverlapping(events)

	assert.Empty(e.t, overlapping)
	assert.Error(e.t, err)
	assert.True(e.t, strings.Contains(err.Error(), "has to be greater than 1."))
}

func (e *EventServiceTestSuite) TestShouldReturnOverlappingEventsWhenOverlapping() {
	events := domain.CreateFakeEvents(2)
	events[1].StartDate = events[0].StartDate.Add(time.Minute * -60)
	events[1].EndDate = events[0].EndDate.Add(time.Minute * 120)

	overlapping, err := EventService{}.AllEventsOverlapping(events)

	assert.NotNil(e.t, overlapping)
	assert.Nil(e.t, err)
	assert.Len(e.t, overlapping, 1)
}

func (e *EventServiceTestSuite) TestShouldReturnAllOverlappingEventsWhenOverlapping() {
	events := domain.CreateFakeEvents(3)
	events[1].StartDate = events[0].StartDate.Add(time.Minute * -60)
	events[1].EndDate = events[0].EndDate.Add(time.Minute * 120)
	events[2].StartDate = events[0].StartDate.Add(time.Minute * -120)
	events[2].EndDate = events[0].EndDate.Add(time.Minute * 190)

	overlapping, err := EventService{}.AllEventsOverlapping(events)

	assert.NotNil(e.t, overlapping)
	assert.Nil(e.t, err)
	assert.Len(e.t, overlapping, 3)
}

func (e *EventServiceTestSuite) TestShouldReturnOvelapingWhenEventsHasSameDate() {
	events := domain.CreateFakeEvents(2)
	events[1].StartDate = events[0].StartDate
	events[1].EndDate = events[0].EndDate

	overlapping, err := EventService{}.AllEventsOverlapping(events)

	assert.NotNil(e.t, overlapping)
	assert.Nil(e.t, err)
	assert.Len(e.t, overlapping, 1)
}

func (e *EventServiceTestSuite) TestShouldReturnErrorWhenEventsInvalid() {
	events := []domain.Event{domain.Event{}, domain.Event{}}

	_, err := EventService{}.IsEventsValid(events)

	assert.Error(e.t, err)
	assert.True(e.t, strings.Contains(err.Error(), "Event.Name"))
	assert.True(e.t, strings.Contains(err.Error(), "Event.StartDate"))
	assert.True(e.t, strings.Contains(err.Error(), "Event.EndDate"))
}
