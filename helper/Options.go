package helper

import (
	"errors"

	"github.com/charmbracelet/bubbles/list"
)

var options = map[string]int{
	"Add a todo":     1,
	"Show all todos": 2,
}

func GetOptions() []list.Item {
	var items []list.Item
	for key := range options {
		items = append(items, item(key))
	}
	return items
}

// GetOptionKey returns the integer key for the specified option
func GetOptionKey(option string) (int, error) {
	key, ok := options[option]
	if !ok {
		return 0, errors.New("Invalid option")
	}
	return key, nil
}
