package timer

import (
	"strconv"

	"github.com/davidv171/timer/figlet"
)

func Start(limit *int) {

	for i := 0; i <= *limit; i++ {
		cmd := strconv.Itoa(i)
		figlet.Figlet(&cmd)

	}
}
