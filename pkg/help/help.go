package help

import (
	"log"

	"github.com/odilonjk/gofolio/pkg/command"
)

type help struct{}

// New help command
func New() command.Cmd {
	return help{}
}

// Execute help command, printing all available commands
func (h help) Execute() {
	log.Println("These are the available commands:")
	log.Println("'ginv <view>' (TODO)")
	log.Println("'ginv buy <ticket> <volume> <price> <currency> <broker>'")
	log.Println("'ginv sell <ticket> <volume> <price> <currency> <broker>'")
}
