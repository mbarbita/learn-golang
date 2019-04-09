package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {

	type CmdArgs struct {
		host string
	}

	var cmdArgs = new(CmdArgs)
	flag.StringVar(&cmdArgs.host, "host", "192.168.0.1", "host addr")
	flag.Parse()

	i := 1
	for {
		cmdOut, err := exec.Command("tvnserver", "-controlservice", "-connect", cmdArgs.host).Output()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			// os.Exit(1)
		}
		fmt.Println(i, cmdOut)
		i++
		time.Sleep(time.Minute)
	}
}
