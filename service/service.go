package service

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/grandcat/zeroconf"
)

type Service struct {
	name     string
	service  string
	domain   string
	host     string
	ip       string
	port     int
	waitTime int
}

func New() *Service {
	return &Service{}
}

type Config struct {
	Name     string `yaml:"Name"`
	Service  string `yaml:"Service"`
	Domain   string `yaml:"Domain"`
	Host     string `yaml:"Host"`
	Ip       string `yaml:"Ip"`
	Port     int    `yaml:"Port"`
	WaitTime int    `yaml:"WaitTime"`
}

func (p *Service) Run(config *Config) error {
	p.name = config.Name
	p.service = config.Service
	p.domain = config.Domain
	p.host = config.Host
	p.ip = config.Ip
	p.port = config.Port
	p.waitTime = config.WaitTime
	// p = &config
	server, err := zeroconf.RegisterProxy(p.name, p.service, p.domain, p.port, p.host, []string{p.ip}, []string{"txtv=0", "lo=1", "la=2"}, nil)
	if err != nil {
		panic(err)
	}
	defer server.Shutdown()
	log.Println("Published proxy service:")
	log.Println("- Name:", p.name)
	log.Println("- Type:", p.service)
	log.Println("- Domain:", p.domain)
	log.Println("- Port:", p.port)
	log.Println("- Host:", p.host)
	log.Println("- IP:", p.ip)

	// Clean exit.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	// Timeout timer.
	// var tc <-chan time.Time
	// if *waitTime > 0 {
	// 	tc = time.After(time.Second * time.Duration(*waitTime))
	// }

	select {
	case <-sig:
		// Exit by user
		// case <-tc:
		// Exit by timeout
	}

	log.Println("Shutting down.")
	return nil
}
