package command

import (
	"log"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/tidwall/buntdb"
)

type cents interface{}

// Buy keeps the buy order data
type Buy struct {
	Ticket, Currency, Broker string
	Volume                   int
	Price                    int64
}

// NewBuy buy command
func NewBuy(args []string) Buy {
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

	return Buy{
		Ticket:   args[0],
		Volume:   v,
		Price:    p,
		Currency: args[3],
		Broker:   strings.Join(args[4:], " "),
	}
}

// Execute the buy order for the ticket
func (b Buy) Execute() {
	// TODO - move this db logic to another package
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Println("Error trying to open connection with database.")
		log.Fatalf(err.Error())
	}
	defer db.Close()

	id, err := uuid.NewRandom()
	if err != nil {
		log.Println("Error during buy operation persistence process.")
		log.Fatalf(err.Error())
	}

	err = db.Update(func(tx *buntdb.Tx) error {
		tx.Set("buy:"+id.String()+":ticket", b.Ticket, nil)
		tx.Set("buy:"+id.String()+":volume", strconv.Itoa(b.Volume), nil)
		tx.Set("buy:"+id.String()+":price", strconv.FormatInt(b.Price, 10), nil)
		tx.Set("buy:"+id.String()+":currency", b.Currency, nil)
		tx.Set("buy:"+id.String()+":broker", b.Broker, nil)

		// TODO - get the ticket from the wallet (if exists) and update it

		return nil
	})
	log.Println("Buy order saved.")
}
