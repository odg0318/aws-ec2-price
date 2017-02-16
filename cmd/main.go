package main

import (
	"fmt"
	"os"

	"github.com/odg0318/aws-ec2-price/pkg/rest"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ec2-instance-price"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   "port",
			EnvVar: "PORT",
			Value:  8080,
		},
	}
	app.Action = func(c *cli.Context) error {
		port := c.Int("port")

		r := rest.GetRouter()
		r.Run(fmt.Sprintf(":%d", port))
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
