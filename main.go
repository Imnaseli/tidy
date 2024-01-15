package main

import (
	"log"
	"os"
)

func main() {
	app := InitApplication()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("Error starting tidy: ", err)
	}
}
