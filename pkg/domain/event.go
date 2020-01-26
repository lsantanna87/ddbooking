package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type Event struct {
	Name        string    `json:"name" validate:"required"`
	StartDate   time.Time `json:"start_date" validate:"required"`
	EndDate     time.Time `json:"end_date" validate:"required"`
	Description string    `json:"description"`
}

type EventsOverlaping struct {
	FirstEvent  Event
	SecondEvent Event
}

func (e Event) IsValid() (bool, error) {
	validate := validator.New()

	if err := validate.Struct(e); err != nil {
		return false, err.(validator.ValidationErrors)
	}

	if ok, err := e.isStartDateBeforeEndDate(); !ok {
		return false, err
	}

	return true, nil
}

func (e Event) isStartDateBeforeEndDate() (bool, error) {
	if !e.StartDate.Before(e.EndDate) {
		erroMsg := fmt.Sprintf("start_date is after end_date for event %s", e.Description)
		return false, errors.New(erroMsg)
	}
	return true, nil
}
