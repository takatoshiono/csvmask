package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/takatoshiono/csvmask/mask"
)

const (
	exitOK = iota
	exitErr
)

var (
	showHelp   bool
	ruleStr    string
	skipHeader bool
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	flag.BoolVar(&showHelp, "help", false, "Show help")
	flag.StringVar(&ruleStr, "rule", "", "Comma separated masking rule of csv")
	flag.BoolVar(&skipHeader, "skipheader", false, "Skip first line of file as header")
	flag.Parse()

	if showHelp {
		flag.PrintDefaults()
		return exitOK
	}

	if ruleStr == "" {
		flag.PrintDefaults()
		return exitErr
	}

	rules, err := mask.NewRules(strings.Split(ruleStr, ","))
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to new rules: %v\n", err)
		return exitErr
	}
	reader := mask.NewReader(os.Stdin, rules)
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
