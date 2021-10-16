package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"os"
)

// errors
var (
	errEmptyInput = errors.New("empty data")
)

// mappingHeaders mapping first line on csv
func mappingHeaders(rows [][]string) (map[int]string, error) {

	// header
	line := rows[0]

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
func generateJson(rows [][]string) (string, error) {

	headers, err := mappingHeaders(rows)

	if err != nil {
		return "", err
	}

	values := make([]map[string]string, 0)

	for idy, line := range rows {

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

// readCsvFromStdin
func readCsvFromStdin(fDelimiter string) ([][]string, error) {

	// reading from stdin if not provide any path in argument
	r := csv.NewReader(os.Stdin)

	r.Comma = rune(fDelimiter[0])
	// TODO: customize comment on csv file
	r.Comment = '#'

	return r.ReadAll()
}
