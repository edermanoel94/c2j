package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"io"
)

var (
	errFormatCsvFile = errors.New("empty file or not format withj csv")
)

func mappingWithHeaders(rows [][]string) map[int]string {

	headerKeys := make(map[int]string)

	row := rows[0]

	for idx, column := range row {
		headerKeys[idx] = column
	}

	return headerKeys

}

func toJson(headerKeys map[int]string, rows [][]string) ([]byte, error) {

	values := make([]map[string]string, 0)

	for _, line := range rows {

		value := make(map[string]string)

		for idx, column := range line {
			value[headerKeys[idx]] = column
		}

		values = append(values, value)
	}

	return json.Marshal(&values)
}

func convert(reader io.Reader, delimiter string, noHeader bool) (string, error) {

	rows, err := csvFromReader(reader, delimiter)

	if err != nil {
		return "", err
	}

	if noHeader {

		rows = rows[1:]

		jsonBytes, err := json.Marshal(rows)

		if err != nil {

			return "", err
		}

		return string(jsonBytes), nil
	}

	var headerKeys map[int]string

	headerKeys = mappingWithHeaders(rows)

	rows = rows[1:]

	jsonBytes, err := toJson(headerKeys, rows)

	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}

func csvFromReader(reader io.Reader, delimiter string) ([][]string, error) {

	//TODO: reading from stdin if not provide any path in argument
	r := csv.NewReader(reader)

	r.Comma = rune(delimiter[0])
	// TODO: customize comment on csv file
	r.Comment = '#'

	rows, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	if len(rows) < 1 {
		return nil, errFormatCsvFile
	}

	return rows, nil
}
