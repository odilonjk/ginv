package main

import (
	"log"
	"os"

	"github.com/odilonjk/gofolio/pkg/buy"
	"github.com/odilonjk/gofolio/pkg/command"
	"github.com/odilonjk/gofolio/pkg/help"
	"github.com/odilonjk/gofolio/pkg/view"
)

func main() {

	if len(os.Args) == 1 {
		log.Println("You must specify what command you want to execute.")
		log.Fatal("Try 'ginv help'")
	}
	typeStr := os.Args[1]

	var c command.Cmd
	if "view" == typeStr {
		c = view.New(os.Args[2:])
	} else if "buy" == typeStr {
		c = buy.New(os.Args[2:])
	} else if "help" == typeStr {
		c = help.New()
	}

	c.Execute()

}
