package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func Execute() {
	app := &cli.App{
		Name:     "Search for Events Overlaping",
		Flags:    CreateFlags(CreateFileFlag, CreateTextFlag),
		Commands: createCommands(createImportCMD, createValidateCMD),
		Before:   validateCLI,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func validateCLI(c *cli.Context) error {
	if c == nil { return fmt.Errorf("context cannot be nil.") }

	if len(c.FlagNames()) > 1 {
		return fmt.Errorf("only one flag is allowed.")
	}

	return nil
}
