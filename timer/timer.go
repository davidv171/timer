package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func startt(timer *int, lolcat *bool) {

	pipecmd := " | figlet -c -k"

	if *lolcat {
		_, err := exec.Command("/bin/bash", "which", "lolcat").Output()

		if err != nil {
			fmt.Println("Is lolcat installed?")
		} else {

			pipecmd = " | lolcat"

		}

	}

	for ; *timer != -1; *timer-- {
		fmt.Println("\033c")

		cmd := "echo " + strconv.Itoa(*timer) + pipecmd

		output, err := exec.Command("/bin/bash", "-c", cmd).Output()

		if err != nil {
			fmt.Println("Error, is figlet installed?")
			os.Exit(1)
		}
		fmt.Println(string(output))
		time.Sleep(1 * time.Second)

	}

}
