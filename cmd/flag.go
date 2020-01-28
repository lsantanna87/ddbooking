package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/lsantanna87/ddbooking/pkg/domain"
	"github.com/pkg/errors"
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

func createInputFromFlags(flag string, ctx *cli.Context) ([]domain.Event, error) {
	process := map[string]func(input string) ([]domain.Event, error){
		"file": processFile,
		"text": processText,
	}

	processFunc, exist := process[flag]
	if !exist {
		return nil, fmt.Errorf("error while invoke process input, flag not found!")
	}

	fmt.Println(ctx.String(flag))
	return processFunc(ctx.String(flag))
}

func processFile(filePath string) ([]domain.Event, error) {
	dat, err := ioutil.ReadFile(filePath)
	if err != nil {
		return []domain.Event{}, errors.Wrap(err, "error when trying to read json file.")
	}

	return domain.Event{}.ToEvents(dat)
}

func processText(textJson string) ([]domain.Event, error) {
	return domain.Event{}.ToEvents([]byte(textJson))
}
