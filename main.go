package main

import (
	"flag"

	"github.com/davidv171/timer/countdown"
	"github.com/davidv171/timer/custom"
)

func main() {

	countdownptr := flag.Int("timer", -1, "Set timer value")
	customf := flag.String("custom", "", "Run custom command after it's over")
	lolcat := flag.Bool("lolcat", false, "Pipe to lolcat?")
	timerptr := flag.Bool("timer", false, "Perform a timer instead of a countdown")

	flag.Parse()

	if *countdownptr == -1 {

	}
	countdown.Countdown(countdownptr, lolcat)
	custom.Runc(customf)

}
