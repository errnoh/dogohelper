// +build darwin freebsd netbsd openbsd
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func openURL(link string) {
	cmd := exec.Command("open", link)
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
