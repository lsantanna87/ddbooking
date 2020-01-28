package domain

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

type Event struct {
	Name      string    `json:"name" validate:"required"`
	StartDate time.Time `json:"start_date" validate:"required"`
	EndDate   time.Time `json:"end_date" validate:"required"`
}

type EventsOverlaping struct {
	FirstEvent  Event
	SecondEvent Event
}

func (e Event) IsValid() (bool, error) {
	validate := validator.New()

	if err := validate.Struct(e); err != nil {
		return false, errors.Wrap(err.(validator.ValidationErrors), fmt.Sprintf("Event: %+v", e))
	}

	if ok, err := e.isStartDateBeforeEndDate(); !ok {
		return false, err
	}

	return true, nil
}

func (e Event) isStartDateBeforeEndDate() (bool, error) {
	if !e.StartDate.Before(e.EndDate) {
		return false, fmt.Errorf("start_date %v is after end_date %v for event", e.StartDate, e.EndDate)
	}

	return true, nil
}

func (e Event) ToEvents(b []byte) ([]Event, error) {
	var events []Event

	if err := json.Unmarshal(b, &events); err != nil {
		return events, errors.Wrap(err, "error when trying to serialize Events from []byte")
	}

	return events, nil
}
