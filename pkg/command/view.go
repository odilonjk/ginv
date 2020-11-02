package command

import (
	"log"
	"strings"

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
	db.CreateIndex("buys", "buy:*", buntdb.IndexString)

	err = db.View(func(tx *buntdb.Tx) error {
		i := investment{}
		err := tx.Ascend("buys", func(k, v string) bool {
			//FIXME - It's not working. It needs a map to work on multiple investments and fill based on key ID

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
				log.Printf("Investiment: %v", i)
				i = investment{}
			}
			log.Printf("key: %s - value: %s", k, v)
			return true
		})
		return err
	})
	if err != nil {
		log.Println("Error while reading records from database.")
		log.Fatalf(err.Error())
	}
}

func (i investment) isFilled() bool {
	return i.ticket != "" && i.broker != "" && i.currency != "" && i.price != "" && i.volume != ""
}
