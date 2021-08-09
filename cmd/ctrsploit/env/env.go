package env

import (
	"github.com/ctrsploit/ctrsploit/env"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name: env.SubCommandName,
	//Aliases: []string{"e"},
	Usage: "gather information",
	Subcommands: []*cli.Command{
		whereCommand,
		graphdriverCommand,
		cgroupsCommand,
		capabilityCommand,
		seccompCommand,
		apparmorCommand,
		fdiskCommand,
	},
}
