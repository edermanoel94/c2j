package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const (
	CmdLongDelimiter  = "--delimiter"
	CmdShortDelimiter = "-d"
)

var usage string = `
USAGE:
  c2j [OPTION...]
  c2j --delimiter -d <string>      choose a delimiter for csv

EXAMPLE:
  cat example.csv | c2j --delimiter ";"
`

var (
	errNotProviderArgs = errors.New("not provided any arguments")
	errCommandNotFound = errors.New("command not found")
)

type option struct {
	delimiter string
	comment   string
	header    bool
}

func parse(opt *option, args ...string) error {

	if len(args) < 1 {
		return nil
	}

	switch args[0] {
	case CmdLongDelimiter,
		CmdShortDelimiter:
		opt.delimiter = args[1]
	default:
		return errCommandNotFound
	}

	return nil
}

func main() {

	args := os.Args[1:]

	opt := option{
		delimiter: ",",
		comment:   "#",
	}

	err := parse(&opt, args...)

	must(err)

	// reading from stdin
	r := csv.NewReader(os.Stdin)

	r.Comma = rune(opt.delimiter[0])
	r.Comment = rune(opt.comment[0])

	lines, err := r.ReadAll()

	must(err)

	if len(lines) < 1 {
		must(errors.New("dont have any lines"))
	}

	jsonValue, err := generateJson(lines)

	must(err)

	fmt.Fprint(os.Stdout, jsonValue)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// mappingHeaders mapping first line on csv
func mappingHeaders(lines [][]string) (map[int]string, error) {

	// header
	line := lines[0]

	if len(line) < 1 {
		return nil, errors.New("dont have any columns")
	}

	headers := make(map[int]string)

	for idx, column := range line {
		headers[idx] = column
	}

	return headers, nil
}

// generateJson
func generateJson(lines [][]string) (string, error) {

	headers, err := mappingHeaders(lines)

	if err != nil {
		return "", err
	}

	values := make([]map[string]string, 0)

	for idy, line := range lines {

		// ignore headers if flag is activated
		if idy == 0 {
			continue
		}

		value := make(map[string]string)

		for idx, column := range line {
			value[headers[idx]] = column
		}

		values = append(values, value)
	}

	jsonValues, err := json.Marshal(&values)

	if err != nil {
		return "", err
	}

	return string(jsonValues), nil
}
