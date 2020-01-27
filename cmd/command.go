package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/lsantanna87/ddbooking/pkg/domain"
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
		Action: processImport,
	}
}

func CreateValidateCMD() *cli.Command {
	return &cli.Command{
		Name:   "validate",
		Usage:  "Validate if events are valid",
		Action: processValidate,
	}
}

func processValidate(c *cli.Context) error {
	if len(c.FlagNames()) > 1 { // it will return
		return fmt.Errorf("only one flag is allowed.")
	}

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

func processImport(c *cli.Context) error {
	if len(c.FlagNames()) > 1 { // it will return
		return fmt.Errorf("only one flag is allowed.")
	}

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

func createInputFromFlags(c *cli.Context) (events []domain.Event, err error) {
	switch c.String(c.FlagNames()[0]) {
	case "text":
		return processText(c.String("text"))
	case "file":
		return processFile(c.String("text"))
	default:
		return
	}
}

func processFile(filePath string) ([]domain.Event, error) {
	file := readJSONFile(filePath)
	return domain.Event{}.ToEvents(file)
}

func processText(textJson string) ([]domain.Event, error) {
	return domain.Event{}.ToEvents([]byte(textJson))
}

func readJSONFile(filePath string) []byte {
	dat, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatalf("error when trying to read json file. %+v", err)
	}

	return dat
}
