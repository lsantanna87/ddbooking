package cmd

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func Execute() {
	app := &cli.App{
		Name:     "Search for Events Overlaping",
		Flags:    CreateFlags(CreateFileFlag, CreateTextFlag),
		Commands: createCommands(createImportCMD, createValidateCMD),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
