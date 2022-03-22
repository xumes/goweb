package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate returns a cli application
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "goweb cli interface"
	app.Usage = "Find IPs and Server names on the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.ca",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Find IPs on thte internet",
			Flags:  flags,
			Action: findIPs,
		},
		{
			Name:   "servers",
			Usage:  "Find the server names who hosts the sites on thte internet",
			Flags:  flags,
			Action: findServers,
		},
	}

	return app
}

func findIPs(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println((ip))
	}
}

func findServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println((server.Host))
	}
}
