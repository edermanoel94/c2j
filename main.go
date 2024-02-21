package main

import (
	"flag"
	"fmt"
	"io"
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
  cat csv_without_header.csv | c2j --no-header | jq`
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
	flag.BoolVar(&fNoHeader, "no-header", false, "parse csv without header, this will use matrices")
	flag.BoolVar(&fNoHeader, "H", false, "parse csv without header")
	flag.BoolVar(&fVersion, "version", false, "print version")
	flag.BoolVar(&fVersion, "v", false, "print version")
	flag.BoolVar(&fHelp, "help", false, "print help")
	flag.BoolVar(&fHelp, "h", false, "print help")

	flag.Usage = func() {
		fmt.Fprintln(os.Stdout, usageString)
		os.Exit(0)
	}
	flag.Parse()

	run(os.Stdin)
}

func run(reader io.Reader) {
	switch {
	case fHelp:
		printUsage()
	case fVersion:
		printVersion()
	case fDelimiter != "":
		output, err := convert(reader, fDelimiter, fNoHeader)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}
		fmt.Fprint(os.Stdout, fmt.Sprintf("%s\n", string(output)))
		os.Exit(0)
	case fDelimiter == "" && flag.NArg() == 0 && (!fHelp || !fVersion):
		output, err := convert(reader, ",", fNoHeader)
		if err != nil {
			fmt.Fprintf(os.Stderr, err.Error())
			os.Exit(-1)
		}
		fmt.Fprint(os.Stdout, fmt.Sprintf("%s\n", string(output)))
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
