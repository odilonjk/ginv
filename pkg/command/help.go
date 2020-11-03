package command

import (
	"log"
)

// Help represents the help command
type Help struct{}

// NewHelp help command
func NewHelp() Help {
	return Help{}
}

// Execute help command, printing all available commands
func (h Help) Execute() {
	log.Println("These are the available commands:")
	log.Println("'ginv <view>' (TODO)")
	log.Println("'ginv buy <ticket> <volume> <price> <currency> <broker>'")
	log.Println("'ginv sell <ticket> <volume> <price> <currency> <broker>'")
}
