package helper 

import "github.com/urfave/cli/v2"

var AllFlags = []cli.Flag{
    &cli.StringFlag{
        Name: "lang",
        Value: "English",
        Usage: "Initial Language",
    },
}
