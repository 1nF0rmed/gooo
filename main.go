package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command struct {
	name    string
	options []string
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	cache := make(map[string]string)

loop:
	for {
		command := getUserCommand(reader)

		if len(command.options) < 1 {
			fmt.Println("Invalid command")
			continue
		}

		switch command.name {
		case "set":
			// set x val
			err := setVal(cache, &command.options[0], &command.options[1])

			if err != nil {
				fmt.Println(err)
			}

		case "get":
			_, exists := cache[command.options[0]]
			if exists {
				fmt.Println("Value: ", cache[command.options[0]])
			} else {
				fmt.Println("Element not present")
			}
		case "delete":
			_, exists := cache[command.options[0]]
			if !exists {
				fmt.Println("Element not present")
				break
			}
			delete(cache, command.options[0])
		case "update":
			_, exists := cache[command.options[0]]
			if !exists {
				fmt.Println("Element not present")
				break
			}
			cache[command.options[0]] = command.options[1]
		case "exit":
			break loop
		}
	}

}

func setVal(cache map[string]string, key *string, value *string) error {
	_, exists := cache[*key]

	if exists {
		return fmt.Errorf("element already present: %s", *key)
	}

	cache[*key] = *value

	return nil
}

func getUserCommand(reader *bufio.Reader) Command {
	fmt.Print("Enter your input: ")
	inputCommand, _ := reader.ReadString('\n')

	inputCommand = strings.TrimSpace(inputCommand)

	delimiter := " "
	commandContent := strings.Split(inputCommand, delimiter)
	return Command{commandContent[0], commandContent[1:]}
}
