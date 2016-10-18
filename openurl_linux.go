// +build linux
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func openURL(link string) {
	cmd := exec.Command("x-www-browser", link)
	err := cmd.Run()
	if err != nil {
		cmd = exec.Command("xdg-open", link)
		err = cmd.Run()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Couldn't open link with x-www-browser or xdg-open")
			return
		}
	}
}
