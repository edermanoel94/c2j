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
  c2j --delimiter -d <string>      choose a delimiter for csv

EXAMPLE:
  cat example.csv | c2j --delimiter ";"
`

var (
	errNotProviderArgs = errors.New("not provided any arguments")
)

type argument struct {
	delimiter string
	//comment   string
}

func parse(arg *argument, args ...string) error {

	if len(args) < 1 {
		return nil
	}

	switch args[0] {
	case CmdLongDelimiter,
		CmdShortDelimiter:
		arg.delimiter = args[1]
	default:
		return errors.New("command not found")
	}

	return nil
}

func main() {

	args := os.Args[1:]

	arg := argument{
		delimiter: ",",
	}

	err := parse(&arg, args...)

	must(err)

	r := csv.NewReader(os.Stdin)

	r.Comma = rune(arg.delimiter[0])
	// r.Comment = rune(arg.comment[0])

	lines, err := r.ReadAll()

	must(err)

	headers := mappingHeaders(lines)

	fmt.Fprint(os.Stdout, generateJson(lines, headers))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func mappingHeaders(lines [][]string) map[int]string {
	headers := make(map[int]string)
	for idy, line := range lines {
		for idx, column := range line {
			if idy == 0 {
				headers[idx] = column
			}
		}
	}
	return headers
}

func generateJson(lines [][]string, headers map[int]string) string {

	values := make([]map[string]string, 0)

	for idy, line := range lines {

		// should ignore headers?
		if idy == 0 {
			continue
		}

		value := make(map[string]string)

		for idx, column := range line {
			value[headers[idx]] = column
		}

		values = append(values, value)
	}

	valuesJson, err := json.Marshal(&values)

	must(err)

	return string(valuesJson)
}
