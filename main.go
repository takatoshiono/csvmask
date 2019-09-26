package main

import (
	"os"
)

var version = "dev"

func main() {
	cli := &cli{in: os.Stdin, outWriter: os.Stdout, errWriter: os.Stderr}
	os.Exit(cli.run(os.Args))
}
