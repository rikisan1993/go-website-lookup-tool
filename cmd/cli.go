package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Website lookup tools"
	app.Usage = "Lookup IP Address, CNAME, MX records, and Name Server for given website"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "Look up the Name Server of a website",
			Flags: flags,
			Action: func(c *cli.Context) error {
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}

				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}

				return nil
			},
		},
		{
			Name:  "ip",
			Usage: "Resolve IP Address of a website",
			Flags: flags,
			Action: func(c *cli.Context) error {
				ip, err := net.LookupIP(c.String("host"))
				if err != nil {
					return err
				}

				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}

				return nil
			},
		},
		{
			Name:  "cname",
			Usage: "Lookup Canonical Name of a website",
			Flags: flags,
			Action: func(c *cli.Context) error {
				cname, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					return err
				}

				fmt.Println(cname)
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Lookup MX Records of a website",
			Flags: flags,
			Action: func(c *cli.Context) error {
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					return err
				}

				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host)
				}
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
