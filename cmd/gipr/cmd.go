package main

import (
	"log"
	"os"

	gipr "github.com/hayashiki/git-issue-pr-release"
	"github.com/urfave/cli"
)

var (
	token = ""
	base  = "" // staging
	head  = "" //"develop"
)

func main() {
	app := cli.NewApp()
	app.Version = "v1.0.0"
	app.Name = "cli tool template"
	app.Usage = "for template"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "token, t",
			EnvVar:      "GITHUB_TOKEN",
			Usage:       "github token",
			Destination: &token,
		},
		cli.StringFlag{
			Name:        "head",
			EnvVar:      "HEAD",
			Usage:       "head",
			Destination: &head,
		},
		cli.StringFlag{
			Name:        "base",
			EnvVar:      "BASE",
			Usage:       "base",
			Destination: &base,
		},
	}

	app.Action = func(c *cli.Context) error {
		gipr.Run(token, head, base)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
