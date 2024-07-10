package helper

import (
	"github.com/urfave/cli/v2"
)

func InitApplication() *cli.App {

	// prepare the databse.

	app := &cli.App{
		Name:  "tidy",
		Usage: "Help organize your life, straight from your terminal.",
		Flags: AllFlags,
		Action: func(c *cli.Context) error {
			if c.NArg() > 0 {
				HandleEvidentArguments(c.Args())
			} else {
				Home()
			}
			return nil
		},
	}

	return app
}

func Home() {
	option := HandleNoArguments() // get the option from the list
	HandleListSelect(option)      // map the option to commands
}
