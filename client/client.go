package client

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/grandcat/zeroconf"

	hostfile "github.com/guumaster/hostctl/pkg/file"
	"github.com/guumaster/hostctl/pkg/types"
)

var (
	mdns_name string = "mdns-resolv"
)

type Client struct {
	hostfile *hostfile.File
	resolver *zeroconf.Resolver
}

func New() *Client {
	return &Client{}
}

type Config struct {
	Service  string `yaml:"Service"`
	Domain   string `yaml:"Domain"`
	WaitTime int    `yaml:"WaitTime"`
}

func (c *Client) Run(config *Config) error {

	h, err := hostfile.NewFile(getDefaultHostFile())
	if err != nil {
		return err
	}
	c.hostfile = h

	// Discover all services on the network (e.g. _workstation._tcp)
	c.resolver, err = zeroconf.NewResolver(nil)
	if err != nil {
		log.Fatalln("Failed to initialize resolver:", err.Error())
	}

	entries := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {

		for entry := range results {
			var ip = ""
			if len(entry.AddrIPv6) > 0 {
				ip = entry.AddrIPv6[0].String()

			} else {
				ip = entry.AddrIPv4[0].String()
			}
			c.hostfile.AddRoute(mdns_name, types.NewRoute(ip, entry.HostName))
			c.hostfile.Flush()

			fmt.Printf("Domains '%s' added.\n", entry.HostName)
		}
	}(entries)

	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		sig := <-sigs
		c.hostfile.RemoveProfile(mdns_name)
		c.hostfile.Flush()
		fmt.Println("signal", sig, "called", ". Terminating...")
		cancel()
	}()

	err = c.resolver.Browse(ctx, config.Service, config.Domain, entries)
	if err != nil {
		log.Fatalln("Failed to browse:", err.Error())
	}

	<-ctx.Done()
	return nil
}

func getDefaultHostFile() string {
	if runtime.GOOS == "linux" {
		return "/etc/hosts" //nolint: goconst
	}

	if runtime.GOOS == "windows" {
		return `C:/Windows/System32/Drivers/etc/hosts`
	}

	return "/etc/hosts"
}
