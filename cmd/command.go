package cmd

import (
	"fmt"

	"github.com/lsantanna87/ddbooking/pkg/service"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func CreateCommands(commands ...func() *cli.Command) []*cli.Command {
	var cmds []*cli.Command
	for _, createCommand := range commands {
		cmds = append(cmds, createCommand())
	}

	return cmds
}

func CreateImportCMD() *cli.Command {
	return &cli.Command{
		Name:   "import",
		Usage:  "Import Events",
		Action: commandImport,
	}
}

func CreateValidateCMD() *cli.Command {
	return &cli.Command{
		Name:   "validate",
		Usage:  "Validate if events are valid",
		Action: commandValidate,
	}
}

func commandValidate(c *cli.Context) error {
	events, err := createInputFromFlags(c)
	if err != nil {
		return errors.Wrap(err, "error executing command validate!")
	}

	isValid, err := service.EventService{}.IsEventsValid(events)
	if err != nil {
		return errors.Wrap(err, "error executing command validate!")
	}

	fmt.Println(isValid)
	return nil
}

func commandImport(c *cli.Context) error {
	events, err := createInputFromFlags(c)

	if err != nil {
		return errors.Wrap(err, "error executing command import!")
	}
	overlapingEvents, err := service.EventService{}.OverlapingEvents(events)

	if err != nil {
		return err
	}

	fmt.Println(overlapingEvents)

	return nil
}
