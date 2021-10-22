package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	version     = "0.1.2"
	usageString = `Usage: c2j [flags]

Flags:
	-h, --help           print help information
	-d, --delimiter      choose delimiter for csv
  -v, --version        print version

Examples:
  cat comma.csv | c2j | jq        
  cat semicolon.csv | c2j -delimiter ";" | jq`
)

// flags
var (
	fDelimiter string
	fVersion   bool
	fHelp      bool
)

func main() {
	flag.StringVar(&fDelimiter, "delimiter", "", "choose a delimiter")
	flag.StringVar(&fDelimiter, "d", "", "choose a delimiter")
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
		// use customize delimiter
		convert(fDelimiter)
		os.Exit(0)
	case fDelimiter == "" && flag.NArg() == 0 && (!fHelp || !fVersion):
		// use standard delimiter, which is comma
		convert(",")
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

// must
func must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
