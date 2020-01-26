package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/lsantanna87/ddbooking/pkg/domain"
	"github.com/lsantanna87/ddbooking/pkg/service"
	"github.com/urfave/cli/v2"
)

func Execute() {
	app := &cli.App{
		Name:     "Events Validator",
		Flags:    createFlags(createFileFlag, createTextFlag),
		Commands: createCommands(createImportCMD, createValidateCMD),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func Stuff(input string) {

	var calendar domain.Calendar
	err := json.Unmarshal([]byte(input), &calendar)

	if err != nil {
		panic(err)
	}

	for _, v := range calendar.Events {
		fmt.Println(v.IsValid())
	}

	cal := domain.Calendar{}
	cal.Events = service.SortEventByStartDate(cal.Events)
	eventsOverlaping := service.AllOverlapingEvents(cal)

	fmt.Println(eventsOverlaping)
}

func createCommands(commands ...func() *cli.Command) []*cli.Command {
	var cmds []*cli.Command
	for _, createCommand := range commands {
		cmds = append(cmds, createCommand())
	}

	return cmds
}

func createImportCMD() *cli.Command {
	return &cli.Command{
		Name:    "import",
		Aliases: []string{"i"},
		Usage:   "Import Events",
		Action: func(c *cli.Context) error {
			return nil
		},
	}
}

func createValidateCMD() *cli.Command {
	return &cli.Command{
		Name:    "validate",
		Aliases: []string{"v"},
		Usage:   "Validate if events are valid",
		Action: func(c *cli.Context) error {
			return nil
		},
	}
}

func createFlags(flags ...func() cli.Flag) []cli.Flag {
	var stringFlags []cli.Flag
	for _, flagFunc := range flags {
		stringFlags = append(stringFlags, flagFunc())
	}

	return stringFlags
}

func createFileFlag() cli.Flag {
	return &cli.StringFlag{
		Name:    "file",
		Aliases: []string{"f"},
		Usage:   "Load Events from json `FILE`",
	}
}

func createTextFlag() cli.Flag {
	return &cli.StringFlag{
		Name:    "text",
		Aliases: []string{"t"},
		Usage:   "Load Events from text in JSON format eg: https://gist.github.com/lsantanna87/5aeb75a0e9affc2eb0cfc8f087acb4da",
	}
}
