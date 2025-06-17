package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ls", "-l")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
	pid := os.Getpid()
	fmt.Println("Current process id:", pid)
}
