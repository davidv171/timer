package main

import (
	"flag"
	"fmt"

	"github.com/davidv171/timer/countdown"
	"github.com/davidv171/timer/custom"
	"github.com/davidv171/timer/timer"
)

func main() {

	countdownptr := flag.Int("countdown", -1, "Set countdown value")
	customf := flag.String("custom", "", "Run custom command after it's over")
	timerptr := flag.Bool("timer", false, "Perform a timer instead of a countdown")
	limitptr := flag.Int("limit", 0, "Upper limit of the timer")

	flag.Parse()

	if *timerptr {

		fmt.Println("Starting timer...")
		timer.Start(limitptr)

	} else if *countdownptr != -1 {

		countdown.Start(countdownptr)

	}
	if *customf != "" {

		custom.Runc(customf)

	}

}
