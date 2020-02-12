package figlet

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func Figlet(cmd *string) {

	pipecmd := "echo \"" + *cmd + "\" | figlet -c -k && echo \"CTRL-c (CONTROL + c) to cancel\""
	out, err := exec.Command("/bin/bash", "-c", pipecmd).Output()


	fmt.Println("\033c")

	if err != nil {
		fmt.Println("Error, is figlet installed?")
		os.Exit(1)
	}
	fmt.Println(string(out))
	time.Sleep(1 * time.Second)

}
