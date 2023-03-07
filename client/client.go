package client

import (
	"context"
	"log"
	"time"

	"github.com/grandcat/zeroconf"
)

type Client struct {
	service  string
	domain   string
	waitTime int
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
	// p = &config
	c.domain = config.Domain
	c.service = config.Service
	c.waitTime = config.WaitTime

	// Discover all services on the network (e.g. _workstation._tcp)
	resolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		log.Fatalln("Failed to initialize resolver:", err.Error())
	}

	entries := make(chan *zeroconf.ServiceEntry)
	go func(results <-chan *zeroconf.ServiceEntry) {
		for entry := range results {
			if len(entry.AddrIPv6) > 0 {
				log.Printf("{\"host\": \"%s\",\"ipv6\": \"%s\",\"port\": %d}", entry.HostName, entry.AddrIPv6[0], entry.Port)
			} else {
				log.Printf("{\"host\": \"%s\",\"ipv4\": \"%s\",\"port\": %d}", entry.HostName, entry.AddrIPv4[0], entry.Port)
			}
		}
		// log.Println("No more entries.")
	}(entries)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.waitTime))
	defer cancel()
	err = resolver.Browse(ctx, c.service, c.domain, entries)
	if err != nil {
		log.Fatalln("Failed to browse:", err.Error())
	}

	<-ctx.Done()
	// Wait some additional time to see debug messages on go routine shutdown.
	// time.Sleep(1 * time.Second)
	return nil
}
