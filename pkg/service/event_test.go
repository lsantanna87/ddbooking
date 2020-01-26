package service

import (
	"fmt"
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

func (e *EventServiceTestSuite) TestShouldReturnTrueWhenEventsOverlaping() {
	event1 := domain.CreateFakeEvent()
	event2 := domain.CreateFakeEvent()
	event2.StartDate = event1.StartDate.Add(time.Minute * -10)
	event2.EndDate = event1.EndDate

	isOverlaping := EventService{}.isEventsOverlaping(event1, event2)

	assert.True(e.t, isOverlaping)
}

func (e *EventServiceTestSuite) TestShouldReturnFalseWhenEventsNotOverlaping() {
	event1 := domain.CreateFakeEvent()
	event2 := domain.CreateFakeEvent()
	event2.StartDate = event1.StartDate.Add(time.Minute * 60)
	event2.EndDate = event1.EndDate.Add(time.Minute * 120)

	isOverlaping := EventService{}.isEventsOverlaping(event1, event2)

	assert.False(e.t, isOverlaping)
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

func (e *EventServiceTestSuite) TestShouldReturnErrorWhenCheckForOverlapingEventsWithSingleEvent() {
	events := domain.CreateFakeEvents(1)

	overlaping, err := EventService{}.OverlapingEvents(events)

	assert.Nil(e.t, overlaping)
	assert.Error(e.t, err)
	assert.EqualError(e.t, err, "number of events has to be greater than 1.")
}

func (e *EventServiceTestSuite) TestShouldReturnErrorWhenEventsInvalid() {
	events := domain.CreateFakeEvents(2)
	events[1].Name = ""
	events[1].StartDate = time.Time{}
	events[1].EndDate = time.Time{}

	overlaping, err := EventService{}.OverlapingEvents(events)

	assert.Nil(e.t, overlaping)
	assert.Error(e.t, err)
	fmt.Println(err)
	assert.True(e.t, strings.Contains(err.Error(), "Event.Name"))
	assert.True(e.t, strings.Contains(err.Error(), "Event.StartDate"))
	assert.True(e.t, strings.Contains(err.Error(), "Event.EndDate"))
}

func (e *EventServiceTestSuite) TestShouldReturnOverlapingEventsWhenOverlaping() {
	events := domain.CreateFakeEvents(2)
	events[1].StartDate = events[0].StartDate.Add(time.Minute * -60)
	events[1].EndDate = events[0].EndDate.Add(time.Minute * 120)

	overlaping, err := EventService{}.OverlapingEvents(events)

	assert.NotNil(e.t, overlaping)
	assert.Nil(e.t, err)
	assert.Len(e.t, overlaping, 1)
}
