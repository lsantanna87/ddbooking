package faker

import (
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/lsantanna87/ddbooking/pkg/domain"
)

func CreateEvents(numOfEvents int) (events []domain.Event) {
	gofakeit.Seed(0)

	for index := 0; index < numOfEvents; index++ {
		date := gofakeit.Date()
		newEvent := domain.Event{Name: gofakeit.BeerStyle(), StartDate: date, EndDate: date.Add(time.Minute * 10), Description: gofakeit.BS()}
		events = append(events, newEvent)
	}

	return
}
