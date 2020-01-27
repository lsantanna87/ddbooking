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

// if len(c.FlagNames()) > 1 { // it will return
// return fmt.Errorf("only one flag is allowed.")
// }

func createInputFromFlags(c *cli.Context) (events []domain.Event, err error) {
	flag := c.FlagNames()[0]
	switch flag {
	case "text":
		return processText(c.String("text"))
	case "file":
		return processFile(c.String("text"))
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
