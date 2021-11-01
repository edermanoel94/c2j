package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	version     = "0.1.6\n"
	usageString = `Usage: c2j [flags]

Flags:
	-h, --help           print help information
  -v, --version        print version
	-d, --delimiter      choose delimiter for csv
	-H, --no-header      parse csv without header, generating a key based on indexes

Examples:
  cat comma.csv              | c2j | jq        
  cat semicolon.csv          | c2j --delimiter ";" | jq
  cat csv_without_header.csv | c2j --no-header | jq

`
)

// flags
var (
	fDelimiter string
	fNoHeader  bool
	fVersion   bool
	fHelp      bool
)

func main() {
	flag.StringVar(&fDelimiter, "delimiter", "", "choose a delimiter")
	flag.StringVar(&fDelimiter, "d", "", "choose a delimiter")
	flag.BoolVar(&fNoHeader, "no-header", false, "parse csv without header")
	flag.BoolVar(&fNoHeader, "H", false, "parse csv without header")
	flag.BoolVar(&fVersion, "version", false, "print version")
	flag.BoolVar(&fVersion, "v", false, "print version")
	flag.BoolVar(&fHelp, "help", false, "print help")
	flag.BoolVar(&fHelp, "h", false, "print help")

	flag.Usage = func() {
		fmt.Fprintln(os.Stdout, usageString)
		os.Exit(1)
	}
	flag.Parse()

	run()
}

func run() {
	switch {
	case fHelp:
		printUsage()
		os.Exit(0)
	case fVersion:
		printVersion()
		os.Exit(0)
	case fDelimiter != "":
		if err := convert(os.Stdin, fDelimiter, fNoHeader); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}
		os.Exit(0)
	case fDelimiter == "" && flag.NArg() == 0 && (!fHelp || !fVersion):
		if err := convert(os.Stdin, ",", fNoHeader); err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stdout, "flag provided but not defined %s \n", flag.Args()[0])
		printUsage()
		os.Exit(-1)
	}
}

func printVersion() {
	fmt.Fprintf(os.Stdout, version)
}

func printUsage() {
	fmt.Fprintf(os.Stdout, usageString)
}
