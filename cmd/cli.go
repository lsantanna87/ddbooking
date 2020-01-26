package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/lsantanna87/ddbooking/pkg/domain"
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

func createCommands(commands ...func() *cli.Command) []*cli.Command {
	var cmds []*cli.Command
	for _, createCommand := range commands {
		cmds = append(cmds, createCommand())
	}

	return cmds
}

func SerializeStringJson(input string) domain.Calendar {
	var calendar domain.Calendar
	err := json.Unmarshal([]byte(input), &calendar)

	if err != nil {
		panic(err)
	}

	return calendar
}

func SerializeFile(filePath string) domain.Calendar {
	dat, err := ioutil.ReadFile(filePath)

	if err != nil {
		panic(err)
	}

	var calendar domain.Calendar
	if err := json.Unmarshal(dat, &calendar); err != nil {
		panic(err)
	}

	return calendar
}

func createImportCMD() *cli.Command {
	return &cli.Command{
		Name:    "import",
		Aliases: []string{"i"},
		Usage:   "Import Events",
		Action: func(c *cli.Context) error {
			filePath := c.String("file")
			stringJson := c.String("text")

			if filePath != "" {
				SerializeFile(filePath)
			}
			if stringJson != "" {
				SerializeStringJson(stringJson)
			}

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
