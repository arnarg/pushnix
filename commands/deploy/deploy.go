package deploy

import (
	"github.com/urfave/cli/v2"

	"github.com/arnarg/pushnix/commands/push"
	"github.com/arnarg/pushnix/util"
)

var logger = util.Logger

var Command cli.Command = cli.Command{
	Name:      "deploy",
	Aliases:   []string{"d"},
	Usage:     "Deploy local configuration to remote host using git and ssh",
	Action:    Run,
	Before:    util.RequiredArgPreFunc,
	ArgsUsage: "<remote>",
	Flags: []cli.Flag{
		&cli.BoolFlag{
			Name:    "force",
			Aliases: []string{"f"},
			Usage:   "Force git push",
		},
		&cli.BoolFlag{
			Name:    "upgrade",
			Aliases: []string{"u"},
			Usage:   "Do a channel upgrade with nixos-rebuild",
		},
	},
}

func Run(c *cli.Context) error {
	upgrade := c.Bool("upgrade")
	nixExtraArgs := util.ParseTerminator(c.Args().Slice())

	host, err := push.RunPush(c)
	if err != nil {
		return err
	}

	logger.Printf("Rebuilding NixOS on host %s@%s...\n", host.User, host.Host)
	return util.SSHNixosRebuild(host, upgrade, nixExtraArgs)
}
