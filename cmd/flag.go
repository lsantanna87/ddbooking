package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/lsantanna87/ddbooking/pkg/domain"
	"github.com/urfave/cli/v2"
)

func CreateFlags(flags ...func() cli.Flag) []cli.Flag {
	var stringFlags []cli.Flag
	for _, flagFunc := range flags {
		stringFlags = append(stringFlags, flagFunc())
	}

	return stringFlags
}

func CreateFileFlag() cli.Flag {
	return &cli.StringFlag{
		Name:  "file",
		Usage: "Load Events from json `FILE`",
	}
}

func CreateTextFlag() cli.Flag {
	return &cli.StringFlag{
		Name:  "text",
		Usage: "Load Events from text in JSON format eg: https://gist.github.com/lsantanna87/5aeb75a0e9affc2eb0cfc8f087acb4da",
	}
}

func createInputFromFlags(c *cli.Context) (events []domain.Event, err error) {
	if c == nil {
		return events, fmt.Errorf("context cannot be nil.")
	}

	if len(c.FlagNames()) > 1 {
		return events, fmt.Errorf("only one flag is allowed.")
	}

	flag := c.FlagNames()[0]
	switch flag {
	case "text":
		return processText(c.String("text"))
	case "file":
		return processFile(c.String("file"))
	default:
		return events, fmt.Errorf("flag not valid! %s", flag)
	}
}

func processFile(filePath string) ([]domain.Event, error) {
	dat, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatalf("error when trying to read json file. %+v", err)
	}
	return domain.Event{}.ToEvents(dat)
}

func processText(textJson string) ([]domain.Event, error) {
	return domain.Event{}.ToEvents([]byte(textJson))
}
