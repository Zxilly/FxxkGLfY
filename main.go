package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
)

func main() {
	app := &cli.App{
		Name:  "FxxkGLfY",
		Usage: "FxxkGLfY is a tool for generating a GLfY record. Current only for JiangXi Universities.",
		Commands: []*cli.Command{
			{
				Name:    "configure",
				Aliases: []string{"cfg"},
				Usage:   "Generating configuration",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "config",
						Aliases: []string{"c"},
						Usage:   "Configuration file path to write, use \"ENV\" to show generated environment variables",
						Value:   "token.json",
					},
				},
				Action: configure,
			},
			{
				Name:  "make",
				Usage: "Make a GLfY record with the configuration file or environment variables",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "config",
						Aliases: []string{"c"},
						Usage:   "Configuration file path to read, use \"ENV\" to read from environment variables",
						Value:   "token.json",
					},
				},
				Action: makeRecord,
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
