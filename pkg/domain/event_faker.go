package domain

import (
	"time"

	"github.com/brianvoe/gofakeit"
)

func CreateFakeEvents(numOfEvents int) (events []Event) {
	gofakeit.Seed(0)

	for index := 0; index < numOfEvents; index++ {
		date := gofakeit.Date()
		newEvent := Event{Name: gofakeit.BeerStyle(), StartDate: date, EndDate: date.Add(time.Minute * 10)}
		events = append(events, newEvent)
	}

	return
}

func CreateFakeEvent() (event Event) {
	return CreateFakeEvents(1)[0]
}
