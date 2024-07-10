package helper

import (
	"fmt"
	"os"

	"github.com/sijirama/tidy/commands"
	"github.com/sijirama/tidy/database"
)

func HandleListSelect(option string) {
	if option == "" {
		return
	}
	key, err := GetOptionKey(option)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch key {
	case 1:
		//CallClear()
		todo := commands.AddTodo()
		database.InsertTodo(DatabaseClient, todo)
		Home()

	case 2:
		CallClear()
		// Code to handle case when key is 2
		todos := database.DisplayTodos(DatabaseClient)
		//fmt.Println(todos)
        commands.DisplayTodos(todos)
		Home()
	default:
		// Code to handle other cases
		fmt.Println("Unknown option selected")
	}
}
