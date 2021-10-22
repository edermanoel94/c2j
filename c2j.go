package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// errors
var (
	errEmptyData = errors.New("empty data")
)

// mappingHeaders mapping first line on csv
func mappingHeaders(defined bool, rows [][]string) map[int]string {

	headerKeys := make(map[int]string)

	if defined {

		row := rows[0]

		for idx, column := range row {
			headerKeys[idx] = column
		}

		return headerKeys
	}

	columnsArr := []string{}

	for i := 1; i <= len(rows[0]); i++ {
		columnsArr = append(columnsArr, fmt.Sprintf("key_%d", i))
	}

	for idx, column := range columnsArr {
		headerKeys[idx] = column
	}

	return headerKeys
}

// toJson
func toJson(headerKeys map[int]string, rows [][]string) ([]byte, error) {

	values := make([]map[string]string, 0)

	for idy, line := range rows {

		// ignore headerKeys if flag is activated
		if idy == 0 {
			continue
		}

		value := make(map[string]string)

		for idx, column := range line {
			value[headerKeys[idx]] = column
		}

		values = append(values, value)
	}

	return json.Marshal(&values)
}

func convert(fDelimiter string) error {

	rows, err := readCsvFromStdin(fDelimiter)

	if err != nil {
		return err
	}

	headerKeys := mappingHeaders(true, rows)

	jsonBytes, err := toJson(headerKeys, rows)

	if err != nil {
		return err
	}

	// NOTE: jq read from stdout, not from stderr
	fmt.Fprint(os.Stdout, string(jsonBytes))

	return nil
}

// readCsvFromStdin
func readCsvFromStdin(delimiter string) ([][]string, error) {

	// reading from stdin if not provide any path in argument
	r := csv.NewReader(os.Stdin)

	r.Comma = rune(delimiter[0])
	// TODO: customize comment on csv file
	r.Comment = '#'

	rows, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	if len(rows) < 1 {
		return nil, errEmptyData
	}

	return rows, nil
}
