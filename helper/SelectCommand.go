package helper

import (
	"fmt"
	"os"
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
	fmt.Println("The option picked is: ", option, " with option ", key)
}
