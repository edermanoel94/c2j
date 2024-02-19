package main

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestMappingWithHeaders(t *testing.T) {

	t.Run("should mapping headers", func(t *testing.T) {

		rows := [][]string{
			{"first_name", "last_name"},
		}

		headers := mappingWithHeaders(rows)

		if headers[0] != "first_name" {
			t.Errorf("want %s, given %s", "first_name", headers[0])
		}

		if headers[1] != "last_name" {
			t.Errorf("want %s, given %s", "last_name", headers[1])
		}
	})

	t.Run("should be empty", func(t *testing.T) {
		rows := [][]string{
			{},
		}

		headers := mappingWithHeaders(rows)

		if len(headers) > 0 {
			t.Errorf("len shouldnt be greater than zero, expected: %d, got: %d", 0, len(headers))
		}
	})
}

func TestToJson(t *testing.T) {

	rows := [][]string{
		{"eder", "costa"},
		{"teste", "teste"},
	}

	headerKeys := map[int]string{
		0: "first_name",
		1: "last_name",
	}

	expected := `[{"first_name":"eder","last_name":"costa"},{"first_name":"teste","last_name":"teste"}]`

	jsonBytes, err := toJson(headerKeys, rows)

	if err != nil {
		t.Error(err)
	}

	if string(jsonBytes) != expected {
		t.Errorf("given %s, want %s", jsonBytes, expected)
	}
}

func TestConvert(t *testing.T) {

	t.Run("convert", func(t *testing.T) {

		csvInput := `
first_name,last_name,phone
Charleen,Roche,253-330-9889
Jenica,Briat,393-963-9525
Julie,Josselsohn,898-929-2639
Maddalena,Bessom,479-862-0782
Collete,Feldklein,143-902-5122
`

		data, err := convert(strings.NewReader(csvInput), ",", false)

		if err != nil {
			t.Errorf("%v", err)
		}

		if data == "" {
			t.Errorf("data couldn't be empty")
		}

		if !json.Valid([]byte(data)) {
			t.Errorf("data is not a valid json")
		}
	})

	t.Run("convert without header", func(t *testing.T) {

		csvInput := `
first_name,last_name,phone
Charleen,Roche,253-330-9889
Jenica,Briat,393-963-9525
Julie,Josselsohn,898-929-2639
Maddalena,Bessom,479-862-0782
Collete,Feldklein,143-902-5122
`

		data, err := convert(strings.NewReader(csvInput), ",", true)

		if err != nil {
			t.Errorf("%v", err)
		}

		if data == "" {
			t.Errorf("data couldn't be empty")
		}

		if !json.Valid([]byte(data)) {
			t.Errorf("data is not a valid json")
		}
	})
}
