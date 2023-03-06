package base

import (
	"github.com/kayuii/mdns-discovery/client"

	"github.com/urfave/cli/v2"
)

var clientFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    Service,
		Aliases: []string{"service", "s"},
		Value:   "discovery-service",
		Usage:   "Set the service type of the new service.",
	},
	&cli.StringFlag{
		Name:    Domain,
		Aliases: []string{"domain", "d"},
		Value:   "local.",
		Usage:   "Set the network domain. Default should be fine.",
	},
	&cli.IntFlag{
		Name:    WaitTime,
		Aliases: []string{"wait", "w"},
		Value:   10,
		Usage:   "Duration in [s] to publish service for.",
	},
}

func clientAction(c *cli.Context) error {
	config := &client.Config{
		Service:  c.String(Service),
		Domain:   c.String(Domain),
		WaitTime: c.Int(WaitTime),
	}
	return client.New().Run(config)
}

func NewClient() *cli.Command {
	return &cli.Command{
		Name:    "Client",
		Aliases: []string{"client"},
		Usage:   "Browse for services in your local network",
		Action:  clientAction,
		Flags:   clientFlags,
	}
}
