package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	exitOK = iota
	exitErr
)

var (
	showHelp    bool
	skipHeader  bool
	templateStr string
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	flag.BoolVar(&showHelp, "help", false, "Show help")
	flag.BoolVar(&skipHeader, "skipheader", false, "Skip first line of file as header")
	flag.StringVar(&templateStr, "template", "", "The template of output")
	flag.Parse()

	if showHelp {
		flag.PrintDefaults()
		return exitOK
	}

	if templateStr == "" {
		flag.PrintDefaults()
		return exitErr
	}

	tmpl, err := NewTemplate(templateStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to new template: %v\n", err)
		return exitErr
	}

	reader := NewReader(os.Stdin, tmpl)
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
		fmt.Fprintln(os.Stdout, s)
	}

	return exitOK
}
