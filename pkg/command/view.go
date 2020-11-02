package command

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/Rhymond/go-money"
	"github.com/tidwall/buntdb"
)

// View is the command to print the current portfolio
type View struct {
	args []string
}

type investment struct {
	ticket   string
	volume   string
	price    string
	currency string
	broker   string
}

// NewViewCmd view command to print the portfolio
func NewViewCmd(args []string) View {
	return View{args}
}

// Execute prints the portfolio
func (v View) Execute() {
	// TODO - move this db logic to another package
	db, err := buntdb.Open("data.db")
	if err != nil {
		log.Println("Error trying to open connection with database.")
		log.Fatalf(err.Error())
	}
	defer db.Close()
	err = db.View(func(tx *buntdb.Tx) error {
		i := investment{}
		log.Println("Current portfolio:")
		log.Println("")
		// TODO - needs the right prefix 'investment' to print the portfolio
		err := tx.AscendKeys("buy:*", func(k, v string) bool {
			if strings.Contains(k, "broker") {
				i.broker = v
			} else if strings.Contains(k, "ticket") {
				i.ticket = v
			} else if strings.Contains(k, "volume") {
				i.volume = v
			} else if strings.Contains(k, "price") {
				i.price = v
			} else if strings.Contains(k, "currency") {
				i.currency = v
			}
			if i.isFilled() {
				log.Printf("%v", i)
				i = investment{}
			}
			return true
		})
		return err
	})
	if err != nil {
		log.Println("Error while reading records from database.")
		log.Fatalf(err.Error())
	}
	log.Println("")
}

func (i investment) isFilled() bool {
	return i.ticket != "" && i.broker != "" && i.currency != "" && i.price != "" && i.volume != ""
}

func (i investment) String() string {
	p, err := strconv.ParseInt(i.price, 10, 64)
	if err != nil {
		log.Fatalln("Error while converting price value.")
	}
	m := money.New(p, i.currency)
	return fmt.Sprintf("Ticket: %s, Price: %s, Volume: %s, Broker: %s",
		i.ticket, m.Display(), i.volume, i.broker)
}
