package main

import (
	"fmt"
	"log"
	"os"

	logger "github.com/sijirama/tidy/Logger"
	"github.com/sijirama/tidy/helper"
)

func main() {
    helper.DatabseInit()
	app := helper.InitApplication()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("Error starting tidy: ", err)
		logger.LogToFile( fmt.Sprintf("Error starting tidy: ", err))
	}
    logger.LogToFile("Tidy is ready")
}
