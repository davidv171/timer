package main

import (
	"fmt"
	"os/exec"
)

func runc(custom *string) {

	if *custom != "" {
		out, err := exec.Command("/bin/bash", "-c", *custom).Output()
		if err != nil {
			fmt.Println("Error running custom command", *custom)
		}
		fmt.Println(string(out))

	}
}
