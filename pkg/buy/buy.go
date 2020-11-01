package buy

import (
	"log"
	"strconv"
	"strings"

	"github.com/Rhymond/go-money"
	"github.com/odilonjk/gofolio/pkg/command"
)

type buy struct {
	Ticket string
	Volume int
	Price  *money.Money
	Broker string
}

// New buy command
func New(args []string) command.Cmd {
	if len(args) < 4 {
		log.Println("You must enter the required arguments.")
		log.Fatal("'ginv buy <ticket> <volume> <price> <currency> <broker>'")
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

	return buy{
		Ticket: args[0],
		Volume: v,
		Price:  money.New(p, args[3]),
		Broker: strings.Join(args[4:], " "),
	}
}

// Execute the buy order for the ticket
func (b buy) Execute() {
	log.Printf("Executing buy order: %s %v", b.Ticket, b.Price.Display())
}
