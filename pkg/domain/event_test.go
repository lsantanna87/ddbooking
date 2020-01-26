package domain

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EventTestSuite struct {
	suite.Suite
	t *testing.T
}

func (e *EventTestSuite) SetupSuite() {
	e.t = e.T()

}

func TestEventSuite(t *testing.T) {
	suite.Run(t, new(EventTestSuite))
}

func (e *EventTestSuite) TestEventIsValid() {
	event := CreateFakeEvent()

	isValid, err := event.IsValid()

	assert.True(e.t, isValid)
	assert.Nil(e.t, err)
}

func (e *EventTestSuite) TestEventWithEndDateBeforeStartDateShouldReturnFalse() {
	event := CreateFakeEvent()
	event.EndDate = event.StartDate.Add(time.Minute * -40)

	isValid, err := event.IsValid()

	assert.False(e.t, isValid)
	assert.NotNil(e.t, err)
}

func (e *EventTestSuite) TestEventWithoutNameShouldBeInvalid() {
	event := CreateFakeEvent()
	event.Name = ""

	isValid, err := event.IsValid()

	assert.False(e.t, isValid)
	assert.NotNil(e.t, err)
}

func (e *EventTestSuite) TestEventWithoutStartDateShouldBeInvalid() {
	event := CreateFakeEvent()
	event.StartDate = time.Time{}

	isValid, err := event.IsValid()

	assert.False(e.t, isValid)
	assert.NotNil(e.t, err)
}

func (e *EventTestSuite) TestEventWithoutEndDateShouldBeInvalid() {
	event := CreateFakeEvent()
	event.EndDate = time.Time{}

	isValid, err := event.IsValid()

	assert.False(e.t, isValid)
	assert.NotNil(e.t, err)
}

func (e *EventTestSuite) TestShouldReturnEventsWhenSerializedWithValidText() {
	events := CreateFakeEvents(4)
	eventByteArray, _ := json.Marshal(events)

	eventsSerialized, err := Event{}.ToEvents(eventByteArray)

	assert.Nil(e.t, err)
	assert.Len(e.t, eventsSerialized, 4)
}

func (e *EventTestSuite) TestShouldReturnErrorWhenSerializedWithInvalidText() {
	events := CreateFakeEvents(4)
	eventByteArray, _ := json.Marshal(events)
	eventByteArray = eventByteArray[1:]

	eventsSerialized, err := Event{}.ToEvents(eventByteArray)

	assert.Nil(e.t, eventsSerialized)
	assert.NotNil(e.t, err)
	assert.True(e.t, strings.Contains(err.Error(), "error when trying to serialize Events from []byte"))
}
