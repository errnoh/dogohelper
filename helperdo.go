// dogohelper - Automatically open IRC links on your default browser.
//
// WARNING: Use at your own risk.
// This should only be run on sandboxed enviroment to prevent people from opening malware etc on your machine.
// Also, probably NSFW.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/hagna/goty"
	"os"
	"strings"
)

var server *string = flag.String("server", "irc.nebula.fi:6667", "Server to connect to in format 'irc.nebula.fi:6667'")
var nick *string = flag.String("nick", "koirabotti", "IRC nick to use")
var channel *string = flag.String("channel", "#tkt-fuksit2012", "Channel to join")

func main() {
	flag.Parse()

	if con, err := goty.Dial(*server, *nick); err != nil {
		fmt.Fprintf(os.Stderr, "sic: %s\n", err)
	} else {
		in := bufio.NewReader(os.Stdin)

		con.Write <- "join " + *channel

		go func() {
			for {
				str, ok := <-con.Read
				if ok == false {
					break
				}

				separated := strings.SplitN(str, ":", 3)
				if len(separated) <= 2 {
					continue
				}

				// Skip system messages etc
				if !strings.Contains(separated[1], "PRIVMSG #") {
                                        fmt.Println(separated[2])
					continue
				}

				// Find actual urls
				msgpart := separated[2]
				msg := strings.Split(msgpart, " ")

				for _, word := range msg {
					if strings.HasPrefix(word, "www.") {
						word = "http://" + word
					}
					if strings.HasPrefix(word, "http://") || strings.HasPrefix(word, "https://") {
						fmt.Println(str)
						go openURL(word)
					}
				}

			}
		}()

		for {
			if _, err := in.ReadString('\n'); err != nil {
				fmt.Fprintf(os.Stderr, "sic: %s\n", err)
				break
			}
		}
		if err := con.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "sic: %s\n", err)
		}
	}
}
