package helper

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Handle entry if arguments are passed directly
func HandleEvidentArguments(args cli.Args) {
	fmt.Println(args)
}
