package main

import (
	"fmt"

	"github.com/lsantanna87/ddbooking/pkg/domain"
	"github.com/lsantanna87/ddbooking/pkg/domain/faker"
	"github.com/lsantanna87/ddbooking/pkg/service"
)

func main() {
	cal := domain.Calendar{}
	cal.Events = faker.CreateEvents(100)
	cal.Events = service.SortEventByStartDate(cal.Events)
	eventsOverlaping := service.AllOverlapingEvents(cal)

	fmt.Println(eventsOverlaping)
}
