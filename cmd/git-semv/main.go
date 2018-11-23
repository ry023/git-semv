package main

import (
	"os"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	cli := &CLI{
		outStream: os.Stdout,
		errStream: os.Stderr,
		Prefix:    "v",
		Minor:     true,
		Pre:       false,
	}
	cli.run(os.Args[1:])
}
