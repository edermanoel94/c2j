package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	version     = "0.1.6\n"
	usageString = `Usage: c2j [flags]

Flags:
	-h, --help           print help information
        -v, --version        print version
	-d, --delimiter      choose delimiter for parse csv
	-H, --no-header      parse csv without header fields, generating a key based on indexes
	-o, --output         save output to a file

Examples:
  cat comma.csv              | c2j
  cat semicolon.csv          | c2j --delimiter ";"
  cat csv_without_header.csv | c2j --no-header
`
)

// flags
var (
	// TODO: check if fDelimiter working with \t
	fDelimiter string
	fNoHeader  bool
	fVersion   bool
	fHelp      bool
	fOutput    string
)

func main() {
	flag.StringVar(&fDelimiter, "delimiter", "", "choose a delimiter")
	flag.StringVar(&fDelimiter, "d", "", "choose a delimiter")
	flag.StringVar(&fOutput, "output", "", "save output to a file")
	flag.StringVar(&fOutput, "o", "", "save output to a file")
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
		if err := convert(os.Stdin, writerDst(fOutput), fDelimiter, fNoHeader); err != nil {
			log.Fatalf("failed to convert with %v", err)
		}
		os.Exit(0)
	case fDelimiter == "" && flag.NArg() == 0 && (!fHelp || !fVersion):
		if err := convert(os.Stdin, writerDst(fOutput), ",", fNoHeader); err != nil {
			log.Fatalf("failed to convert with %v", err)
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

func writerDst(outputFile string) io.Writer {
	if outputFile != "" {
		file, err := os.Create(outputFile)
		if err != nil {
			log.Fatalf("couldnt create file: %v", err)
		}
		return file
	}

	return os.Stdout
}
