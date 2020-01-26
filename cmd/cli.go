package cmd

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func Execute() {
	app := &cli.App{
		Name:     "Events Validator",
		Flags:    CreateFlags(CreateFileFlag, CreateTextFlag),
		Commands: CreateCommands(CreateImportCMD, CreateValidateCMD),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
