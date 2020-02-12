package timer

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func Startt(countdown *int, lolcat *bool) {

	pipecmd := " | figlet -c -k"

	if *lolcat {
		_, err := exec.Command("/bin/bash", "which", "lolcat").Output()

		if err != nil {
			fmt.Println("Is lolcat installed?")
		} else {

			pipecmd = " | lolcat"

		}

	}

	for ; *countdown != -1; *countdown-- {
		fmt.Println("\033c")

		cmd := "echo " + strconv.Itoa(*countdown) + pipecmd

		output, err := exec.Command("/bin/bash", "-c", cmd).Output()

		if err != nil {
			fmt.Println("Error, is figlet installed?")
			os.Exit(1)
		}
		fmt.Println(string(output))
		time.Sleep(1 * time.Second)

	}

}