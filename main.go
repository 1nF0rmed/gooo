package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	/*
		Commands:
		 - set x val
		 - get x
		 - update x new_val
		 - delete x

		Approach:
		 - File store - async
		 - Cache for recent values
		 - bloom filter
	*/
	reader := bufio.NewReader(os.Stdin)

	cache := make(map[string]string)

	for {
		fmt.Print("Enter your input: ")
		inputCommand, _ := reader.ReadString('\n')

		inputCommand = strings.TrimSpace(inputCommand)

		// Parse command into
		delimiter := " "
		commandContent := strings.Split(inputCommand, delimiter)

		if len(commandContent) < 2 {
			fmt.Println("Invalid command")
		}

		switch commandContent[0] {
		case "set":
			// set x val
			_, exists := cache[commandContent[1]]
			if exists {
				fmt.Println("Element already present: ", cache[commandContent[1]])
				break
			}
			cache[commandContent[1]] = commandContent[2]
			fmt.Printf("%T\n", commandContent[2])
		case "get":
			_, exists := cache[commandContent[1]]
			if exists {
				fmt.Println("Value: ", cache[commandContent[1]])
			} else {
				fmt.Println("Element not present")
			}
		case "delete":
			_, exists := cache[commandContent[1]]
			if !exists {
				fmt.Println("Element not present")
				break
			}
			delete(cache, commandContent[1])
		case "update":
			_, exists := cache[commandContent[1]]
			if !exists {
				fmt.Println("Element not present")
				break
			}
			cache[commandContent[1]] = commandContent[2]
		case "exit":
			break
		}
	}

}
