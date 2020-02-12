package countdown

import (
	"strconv"

	"github.com/davidv171/timer/figlet"
)

func Start(countdown *int) {

	for ; *countdown != -1; *countdown-- {
		cmd := strconv.Itoa(*countdown)
		figlet.Figlet(&cmd)

	}

}
