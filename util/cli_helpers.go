package util

import (
	"strings"

	"github.com/urfave/cli/v2"
)

func RequiredArgPreFunc(c *cli.Context) error {
	arg := c.Args().Get(0)
	if arg == "" || strings.HasPrefix(arg, "-") {
		cli.ShowCommandHelpAndExit(c, c.Command.Name, 1)
	}
	return nil
}

func ParseTerminator(args []string) []string {
	for i, a := range args {
		if a == "--" {
			return args[i+1:]
		}
	}

	return []string{}
}
