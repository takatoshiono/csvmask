package main

import (
	"flag"
	"fmt"
	"io"
)

const (
	exitOK = iota
	exitErr
)

type (
	cli struct {
		in                   io.Reader
		outWriter, errWriter io.Writer
	}
)

var (
	showHelp    bool
	showVersion bool
	skipHeader  bool
	templateStr string
)

func (c *cli) run(args []string) int {
	flagset := flag.NewFlagSet("cli", flag.ContinueOnError)
	flagset.SetOutput(c.errWriter)
	flagset.BoolVar(&showHelp, "help", false, "Show help")
	flagset.BoolVar(&showVersion, "version", false, "Show version")
	flagset.BoolVar(&skipHeader, "skipheader", false, "Skip first line of file as header")
	flagset.StringVar(&templateStr, "template", "", "The template of output")

	if err := flagset.Parse(args[1:]); err != nil {
		flag.PrintDefaults()
		return exitErr
	}

	if showHelp {
		flag.PrintDefaults()
		return exitOK
	}

	if showVersion {
		fmt.Println(version)
		return exitOK
	}

	if templateStr == "" {
		flag.PrintDefaults()
		return exitErr
	}

	tmpl, err := NewTemplate(templateStr)
	if err != nil {
		fmt.Fprintf(c.errWriter, "failed to new template: %v\n", err)
		return exitErr
	}

	reader := NewReader(c.in, tmpl)
	reader.SkipHeader = skipHeader

	for {
		s, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(c.errWriter, "failed to read: %v\n", err)
			return exitErr
		}
		fmt.Fprintln(c.outWriter, s)
	}

	return exitOK
}
