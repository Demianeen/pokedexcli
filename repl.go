package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()

		args := getCleanedArgs(userInput)
		if len(args) == 0 {
			continue
		}

		commandName := args[0]
		command, ok := commands[commandName]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}
		err := command.callback(cfg, args[1:]...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func getCleanedArgs(userInput string) []string {
	lowerCaseInput := strings.ToLower(userInput)
	args := strings.Fields(lowerCaseInput)
	return args
}
