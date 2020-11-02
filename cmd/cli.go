package main

import (
	"log"
	"os"

	"github.com/odilonjk/ginv/pkg/command"
)

// Cmd is the command to be executed
type Cmd interface {
	// Execute runs command
	Execute()
}

func main() {

	if len(os.Args) == 1 {
		log.Println("You must specify what command you want to execute.")
		log.Fatal("Try 'ginv help'")
	}

	var c Cmd
	switch os.Args[1] {
	case "view":
		c = command.NewViewCmd(os.Args[2:])
	case "buy":
		c = command.NewBuyCmd(os.Args[2:])
	case "help":
		c = command.NewHelpCmd()
	default:
		log.Fatalf("%s is not a recognized command. Type 'ginv help' for more info.", os.Args[1])
	}

	c.Execute()

}
