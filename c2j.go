package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var (
	errEmptyData = errors.New("empty data")
)

func mappingWithHeaders(rows [][]string) map[int]string {

	headerKeys := make(map[int]string)

	row := rows[0]

	for idx, column := range row {
		headerKeys[idx] = column
	}

	return headerKeys

}

func mappingNoHeaders(rows [][]string) map[int]string {

	headerKeys := make(map[int]string)

	columnsArr := make([]string, 0)

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

	for _, line := range rows {

		value := make(map[string]string)

		for idx, column := range line {
			value[headerKeys[idx]] = column
		}

		values = append(values, value)
	}

	return json.Marshal(&values)
}

func convert(src io.Reader, dst io.Writer, delimiter string, noHeader bool) error {

	rows, err := csvFromReader(src, delimiter)

	if err != nil {
		return err
	}

	var headerKeys map[int]string

	if noHeader {
		headerKeys = mappingNoHeaders(rows)
	} else {
		headerKeys = mappingWithHeaders(rows)
		// ignoring header
		rows = rows[1:]
	}

	jsonBytes, err := toJson(headerKeys, rows)

	if err != nil {
		return err
	}

	// NOTE: jq read from stdout, not from stderr
	fmt.Fprintln(dst, fmt.Sprintf("%s", string(jsonBytes)))

	return nil
}

// csvFromReader
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
		return nil, errEmptyData
	}

	return rows, nil
}
