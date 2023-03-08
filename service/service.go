package service

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/grandcat/zeroconf"
)

type Service struct {
	s *zeroconf.Server
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
	var err error
	p.s, err = zeroconf.RegisterProxy(config.Name, config.Service, config.Domain, config.Port, config.Host, []string{config.Ip}, []string{"txtv=0", "lo=1", "la=2"}, nil)
	if err != nil {
		panic(err)
	}
	defer p.s.Shutdown()
	log.Println("Published proxy service:")
	log.Println("- Name:", config.Name)
	log.Println("- Type:", config.Service)
	log.Println("- Domain:", config.Domain)
	log.Println("- Port:", config.Port)
	log.Println("- Host:", config.Host)
	log.Println("- IP:", config.Ip)

	// Clean exit.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	<-sig

	log.Println("Shutting down.")
	return nil
}
