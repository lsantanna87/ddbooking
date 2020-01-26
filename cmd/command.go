package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/lsantanna87/ddbooking/pkg/api"
	"github.com/lsantanna87/ddbooking/pkg/domain"
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
		Name:    "import",
		Aliases: []string{"i"},
		Usage:   "Import Events",
		Action:  processImport,
	}
}

func CreateValidateCMD() *cli.Command {
	return &cli.Command{
		Name:    "validate",
		Aliases: []string{"v"},
		Usage:   "Validate if events are valid",
		Action:  processValidate,
	}
}

func processValidate(c *cli.Context) error {
	events, err := process(c.String("file"), c.String("text"))
	api := api.EventAPI{Events: events}
	api.IsEventValid()

	if err != nil {
		return errors.Wrap(err, "error executing command validate!")
	}
	return nil
}

func processImport(c *cli.Context) error {
	events, err := process(c.String("file"), c.String("text"))
	if err != nil {
		return errors.Wrap(err, "error executing command import!")
	}

	api := api.EventAPI{Events: events}
	fmt.Println(api.GetOverlapingEvents())
	return nil
}

func process(filePath string, textJson string) ([]domain.Event, error) {
	if filePath != "" {
		file := readJSONFile(filePath)
		return domain.Event{}.ToEvents(file), nil
	} else if textJson != "" {
		return domain.Event{}.ToEvents([]byte(textJson)), nil
	} else {
		return []domain.Event{}, fmt.Errorf("file and text are empty.")
	}
}

func readJSONFile(filePath string) []byte {
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("error when trying to read json file. %+v", err)
	}
	return dat
}
