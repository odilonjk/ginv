package view

import (
	"log"

	"github.com/odilonjk/gofolio/pkg/command"
)

// View is the command to print the current portfolio
type view struct {
	args []string
}

// New view command to print the portfolio
func New(args []string) command.Cmd {
	return view{args}
}

// Execute prints the portfolio
func (v view) Execute() {
	log.Println("I'm sorry! View command has not been implemented yet.")
}
