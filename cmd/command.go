package cmd

import (
	"fmt"

	"github.com/lsantanna87/ddbooking/pkg/service"
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
		Usage:  "Validate if events are valid",
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

	fmt.Println(isValid)

	return nil
}

func commandImport(c *cli.Context) error {
	flag := c.FlagNames()[0]

	events, err := createInputFromFlags(flag, c)
	if err != nil {
		return errors.Wrap(err, "error while invoking createInputFromFlags in commandImport.")
	}

	overlapingEvents, err := service.EventService{}.GetAllEventsOverlaping(events)
	if err != nil {
		return errors.Wrap(err, "error while invoking EventService{}.OverlapingEvents in commandImport.")
	}

	fmt.Println(overlapingEvents)

	return nil
}
