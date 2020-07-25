package push

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/arnarg/pushnix/util"
)

var logger = util.Logger

var Command cli.Command = cli.Command{
	Name:      "push",
	Aliases:   []string{"p"},
	Usage:     "Push local configuration to remote host using git",
	Action:    Run,
	Before:    util.RequiredArgPreFunc,
	ArgsUsage: "<remote>",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "force",
			Aliases: []string{"f"},
			Usage:   "Force git push",
		},
	},
}

func Run(c *cli.Context) error {
	_, err := RunPush(c)
	return err
}

func RunPush(c *cli.Context) (*util.SSHHost, error) {
	r := c.Args().Get(0)
	force := c.Bool("force")

	host, err := util.GetHostFromRemoteName(r)
	if err != nil {
		return nil, err
	}
	if host == nil {
		return nil, fmt.Errorf("Could not find remote %s", r)
	}

	logger.Printf("Pushing configuration to remote %s...\n", r)
	return host, util.GitPush(r, force)
}
