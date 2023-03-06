package main

import (
	"log"
	"os"
	"time"

	mdnscli "github.com/kayuii/mdns-discovery"
	base "github.com/kayuii/mdns-discovery/cli"

	"github.com/urfave/cli/v2"
)

func main() {

	app := cli.NewApp()
	app.Version = mdnscli.Version
	app.Usage = "plotting utility for chia."
	app.Compiled = time.Now()
	app.Authors = []*cli.Author{
		{
			Name:  "Kayuii",
			Email: "577738@qq.com",
		},
	}
	app.Commands = []*cli.Command{
		base.NewService(),
		base.NewClient(),
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
