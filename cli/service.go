package base

import (
	"github.com/kayuii/mdns-discovery/service"
	"github.com/urfave/cli/v2"
)

var serviceFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    Name,
		Aliases: []string{"name", "n"},
		Value:   "mdns-discovery",
		Usage:   "The name for the service.",
	},
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
	&cli.StringFlag{
		Name:    Host,
		Aliases: []string{"host"},
		Value:   "host1",
		Usage:   "Set host name for service.",
	},
	&cli.StringFlag{
		Name:    Ip,
		Aliases: []string{"ip", "i"},
		Value:   "::1",
		Usage:   "Set IP a service should be reachable.",
	},
	&cli.IntFlag{
		Name:    Port,
		Aliases: []string{"port", "p"},
		Value:   42424,
		Hidden:  true,
		Usage:   "Set the port the service is listening to.",
	},
	&cli.IntFlag{
		Name:    WaitTime,
		Aliases: []string{"wait", "w"},
		Value:   10,
		Usage:   "Duration in [s] to publish service for.",
	},
}

func serviceAction(c *cli.Context) error {
	config := &service.Config{
		Name:    c.String(Name),
		Service: c.String(Service),
		Domain:  c.String(Domain),
		Host:    c.String(Host),
		Ip:      c.String(Ip),
		Port:    c.Int(Port),
	}
	return service.New().Run(config)
}

func NewService() *cli.Command {
	return &cli.Command{
		Name:    "Service",
		Aliases: []string{"service"},
		Usage:   "Service Discovery with mDNS",
		Action:  serviceAction,
		Flags:   serviceFlags,
	}
}
