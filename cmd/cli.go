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
	typeStr := os.Args[1]

	var c Cmd
	if "view" == typeStr {
		c = command.NewViewCmd(os.Args[2:])
	} else if "buy" == typeStr {
		c = command.NewBuyCmd(os.Args[2:])
	} else if "help" == typeStr {
		c = command.NewHelpCmd()
	}

	c.Execute()

}
