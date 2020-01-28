package cmd

import (
	"os"
	"strconv"

	"github.com/lsantanna87/ddbooking/pkg/domain"
	"github.com/lsantanna87/ddbooking/pkg/service"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func createCommands(commands ...func() *cli.Command) []*cli.Command {
	var cmds []*cli.Command
	for _, createCommand := range commands {
		cmds = append(cmds, createCommand())
	}

	return cmds
}

func createImportCMD() *cli.Command {
	return &cli.Command{
		Name:   "import",
		Usage:  "Import Events",
		Action: commandImport,
	}
}

func createValidateCMD() *cli.Command {
	return &cli.Command{
		Name:   "validate",
		Usage:  "Validate Events",
		Action: commandValidate,
	}
}

func commandValidate(c *cli.Context) error {
	flag := c.FlagNames()[0]

	events, err := createInputFromFlags(flag, c)
	if err != nil {
		return errors.Wrap(err, "error while invoking createInputFromFlags in commandValidate.")
	}

	isValid, err := service.EventService{}.IsEventsValid(events)
	if err != nil {
		return errors.Wrap(err, "error while invoking EventService{}.IsEventsValid in commandValidate.")
	}

	printIsEventsValid(isValid)
	return nil
}

func commandImport(c *cli.Context) error {
	flag := c.FlagNames()[0]

	events, err := createInputFromFlags(flag, c)
	if err != nil {
		return errors.Wrap(err, "error while invoking createInputFromFlags in commandImport.")
	}

	eventsOverlapping, err := service.EventService{}.AllEventsOverlapping(events)
	if err != nil {
		return errors.Wrap(err, "error while invoking EventService{}.OverlappingEvents in commandImport.")
	}

	printEventsOverlapping(eventsOverlapping)

	return nil
}

func printEventsOverlapping(eventsOverlapping []domain.EventsOverlapping) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Event 1", "Event 2", "End Date Event 1", "Start Date Event 2"})

	for _, v := range eventsOverlapping {
		table.Append([]string{v.FirstEvent.Name, v.SecondEvent.Name, v.FirstEvent.EndDate.String(), v.SecondEvent.StartDate.String()})
	}
	table.Render()
}

func printIsEventsValid(isValid bool) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Are events valid?"})
	table.Append([]string{strconv.FormatBool(isValid)})
	table.Render()
}
