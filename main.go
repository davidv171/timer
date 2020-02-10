package main

import (
	"flag"
	"github.com/davidv171/timer/timer"
	"github.com/davidv171/custom/custom"
)

func main() {

	timer := flag.Int("timer", 180, "Set timer value")
	custom := flag.String("custom", "", "Run custom command after it's over")
	lolcat := flag.Bool("lolcat", false, "Pipe to lolcat?")

	flag.Parse()

	timer.Startt(timer, lolcat)

	custom.runc(custom)

}
