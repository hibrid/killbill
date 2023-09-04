package main

import (
	"fmt"
	"os"
)

// Command defines the structure for commands in the CLI
type Command struct {
	name        string
	description string
	handler     func()
}

// List of available commands
var commands = []Command{
	{
		name:        "hello",
		description: "Prints 'Hello, world!'",
		handler:     helloHandler,
	},
	{
		name:        "exit",
		description: "Exits the program",
		handler:     exitHandler,
	},
}

func main() {
	for {
		displayCommands()
		executeCommand()
	}
}

// displayCommands lists all the available commands
func displayCommands() {
	fmt.Println("Available commands:")
	for i, cmd := range commands {
		fmt.Printf("%d. %s: %s\n", i+1, cmd.name, cmd.description)
	}
}

// executeCommand prompts the user to select a command and then executes it
func executeCommand() {
	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	if choice < 1 || choice > len(commands) {
		fmt.Println("Invalid choice")
		return
	}

	selectedCommand := commands[choice-1]
	selectedCommand.handler()
}

// helloHandler is the function associated with the 'hello' command
func helloHandler() {
	fmt.Println("Hello, world!")
}

// exitHandler exits the program
func exitHandler() {
	fmt.Println("Exiting...")
	os.Exit(0)
}

// To add a new command, simply add a new Command instance to the 'commands' slice
// and define its associated handler function.
