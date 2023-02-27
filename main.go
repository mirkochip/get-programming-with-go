package main

import (
	"fmt"
	"os"

	"github.com/mirkochip/get-programming-with-go/common"
	exps "github.com/mirkochip/get-programming-with-go/experiments"
)

// Definition of command
type cmd struct {
	fn func()
}

// Subcommands playlist
var commandFuncs = map[string]cmd{
	"greetings": {
		fn: exps.Greetings,
	},
	"playground": {
		fn: exps.Playground,
	},
	"malacandra": {
		fn: exps.Malacandra,
	},
	"guess": {
		fn: exps.Guess,
	},
	"random-dates": {
		fn: exps.RandomDates,
	},
	"mars-tickets": {
		fn: exps.MarsTickets,
	},
}

func main() {
	exp := os.Args[1]
	fmt.Printf("Running experiment %s\n", exp)
	common.PrintSymbols("#", 29)
	cmd, ok := commandFuncs[exp]
	if !ok {
		fmt.Printf("Argument [%s] not found into the experiments playlist.\n", exp)
		return
	}
	cmd.fn()
}
