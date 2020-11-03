package command

import (
	"log"
	"strconv"
	"strings"
)

// Sell keeps the sell order data
type Sell struct {
	Ticket, Currency, Broker string
	Volume                   int
	Price                    int64
}

// NewSell sell command
func NewSell(args []string) Sell {
	if len(args) < 4 {
		log.Println("You must enter the required arguments.")
		log.Fatal("'ginv sell <ticket> <volume> <price> <currency> <broker>'")
	}
	v, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatalf("Could not process volume argument: %v.", args[1])
	}
	if v < 1 {
		log.Fatalf("%d is not a valid volume.", v)
	}
	p, err := strconv.ParseInt(args[2], 10, 64)
	if err != nil {
		log.Fatalf("Could not process price argument: %v.", args[2])
	}
	if p < 1 {
		log.Fatalf("%d is not a valid price.", v)
	}

	return Sell{
		Ticket:   args[0],
		Volume:   v,
		Price:    p,
		Currency: args[3],
		Broker:   strings.Join(args[4:], " "),
	}
}

// Execute the sell order for the ticket
func (s Sell) Execute() {
	log.Printf("To be implemented.")
}
