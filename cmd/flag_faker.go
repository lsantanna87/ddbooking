package cmd

import (
	"flag"
	"fmt"

	"github.com/urfave/cli/v2"
)

func CreateFakeContextWithFlag(flagName string) *cli.Context {
	set := flag.NewFlagSet(flagName, 0)
	set.String(flagName, "test", "test")
	_ = set.Parse([]string{fmt.Sprintf("--%s", flagName), "--test"})

	return cli.NewContext(nil, set, nil)
}
