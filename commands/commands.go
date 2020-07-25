package commands

import (
	"github.com/urfave/cli/v2"

	"github.com/arnarg/pushnix/commands/deploy"
	"github.com/arnarg/pushnix/commands/push"
)

func GetCommands() []*cli.Command {
	return []*cli.Command{
		&push.Command,
		&deploy.Command,
	}
}
