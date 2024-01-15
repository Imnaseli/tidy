package main

import (

	//"github.com/sijirama/tidy/commands"
	"github.com/sijirama/tidy/helper"
	"github.com/urfave/cli/v2"
)

func InitApplication() *cli.App {
	app := &cli.App{
		Name:  "tidy",
		Usage: "Help organize your life, straight from your terminal.",
		Flags: helper.AllFlags,
		Action: func(c *cli.Context) error {
			if c.NArg() > 0 {
				helper.HandleEvidentArguments(c.Args())
			} else {
				option := helper.HandleNoArguments() // get the option from the list
                helper.HandleListSelect(option) // map the option to commands
			}
			return nil
		},
	}

	return app
}

