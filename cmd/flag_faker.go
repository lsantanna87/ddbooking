package cmd

import (
	"flag"
	"fmt"

	"github.com/urfave/cli/v2"
)

func CreateFakeContextWithFlag(flagName string, argValue string) *cli.Context {
	set := flag.NewFlagSet(flagName, 0)
	set.String(flagName, "test", "test")
	_ = set.Parse([]string{fmt.Sprintf("--%s", flagName), argValue})

	return cli.NewContext(nil, set, nil)
}

func CreateFakeContextWithTwoFlags(flagName string, flagName2 string, argValue string) *cli.Context {
	set := flag.NewFlagSet(flagName, 0)
	set.String(flagName, "test", "test")
	_ = set.Parse([]string{fmt.Sprintf("--%s", flagName), argValue})

	context := cli.NewContext(nil, set, nil)
	set2 := flag.NewFlagSet(flagName2, 0)
	set2.String(flagName2, "test", "test")
	_ = set2.Parse([]string{fmt.Sprintf("--%s", flagName2), argValue})

	return cli.NewContext(nil, set2, context)
}
