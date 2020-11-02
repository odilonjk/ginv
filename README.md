# ginv

I made this very basic CLI to manage an investment portfolio.
It's a basic CLI where you can register buy and sell orders.

You're welcome to give it a try.

The purpose of this project is to study the usage of Golang [interfaces](https://tour.golang.org/methods/9) and how to implement a CLI using Golang.

## Supported investments

It currently have a basic support on every investiment that requires:
- Ticket or name
- Volume
- Price and currency
- Broker name

You could use it to manage stocks, REITs, etc.

## How does it works?

You just need to type on the terminal the command you want to execute.
Example of a buy order: `ginv buy MSFT 50 22019 US TD Ameritrade`

It adds the ticket to your portfolio using the specified arguments.

You can type `ginv help` to list the available commands.

Ps.: If you're running the binary `ginv` outside the OS PATH, remember that dot-slash will be required to run the file. Example: `./ginv help`.

## How to install it?

You will need to clone this source code or download it.

Once you have the source code on your computer, you will need to build it with running on terminal `make build`.

Finally, I'd recommend you to put the generated binary file on the operational system path.
You could do it moving the binary file to `~/.local/bin` and exporting it with `export PATH=$PATH:~/.local/bin`

## Roadmap 

- Add tests!!!
- Add some way to differ one kind of investment from another.
- Add some embedded database on the project.