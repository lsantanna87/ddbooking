package cmd

import "github.com/urfave/cli/v2"

func CreateFlags(flags ...func() cli.Flag) []cli.Flag {
	var stringFlags []cli.Flag
	for _, flagFunc := range flags {
		stringFlags = append(stringFlags, flagFunc())
	}

	return stringFlags
}

func CreateFileFlag() cli.Flag {
	return &cli.StringFlag{
		Name:    "file",
		Aliases: []string{"f"},
		Usage:   "Load Events from json `FILE`",
	}
}

func CreateTextFlag() cli.Flag {
	return &cli.StringFlag{
		Name:    "text",
		Aliases: []string{"t"},
		Usage:   "Load Events from text in JSON format eg: https://gist.github.com/lsantanna87/5aeb75a0e9affc2eb0cfc8f087acb4da",
	}
}
