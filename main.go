package main

import (
	"flag"
	"github.com/davidv171/timer/custom"
	"github.com/davidv171/timer/timer"
)

func main() {

	timerf := flag.Int("timer", 180, "Set timer value")
	customf := flag.String("custom", "", "Run custom command after it's over")
	lolcat := flag.Bool("lolcat", false, "Pipe to lolcat?")

	flag.Parse()

	timer.Startt(timerf, lolcat)

	custom.Runc(customf)

}
