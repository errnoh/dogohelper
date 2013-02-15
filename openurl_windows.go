package main

import (
	"fmt"
	"os"
	"os/exec"
)

func openlink(link string) {
	cmd := exec.Command(link)
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't open link")
	}
}
