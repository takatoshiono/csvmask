package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/takatoshiono/csvmask/mask"
)

const (
	exitOK = iota
	exitErr
)

var (
	showHelp   bool
	skipHeader bool
	template   string
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	flag.BoolVar(&showHelp, "help", false, "Show help")
	flag.BoolVar(&skipHeader, "skipheader", false, "Skip first line of file as header")
	flag.StringVar(&template, "template", "", "The template of output")
	flag.Parse()

	if showHelp {
		flag.PrintDefaults()
		return exitOK
	}

	template, err := mask.NewTemplate(template)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to new template: %v\n", err)
		return exitErr
	}

	reader := mask.NewReader(os.Stdin, template)
	reader.SkipHeader = skipHeader

	for {
		s, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "failed to read: %v\n", err)
			return exitErr
		}
		fmt.Fprint(os.Stdout, s)
	}

	return exitOK
}
