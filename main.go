package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/arnarg/pushnix/commands"
)

var Version string = "unknown"

func main() {
	app := &cli.App{
		Name:     "pushnix",
		Usage:    "push configuration to a NixOS host",
		Version:  Version,
		Commands: commands.GetCommands(),
	}

	// By default log prints timestamp
	// Here I'm removing that
	log.SetFlags(0)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
