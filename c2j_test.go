package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
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

func TestMappingNoHeaders(t *testing.T) {

	t.Run("should mapping without headers", func(t *testing.T) {
		rows := [][]string{
			{"eder", "manoel"},
			{"something", "joao"},
		}

		expected := map[int]string{
			0: "key_1",
			1: "key_2",
		}

		headers := mappingNoHeaders(rows)

		if !reflect.DeepEqual(headers, expected) {
			t.Errorf("should be deep equal, expected: %v, got: %v", expected, headers)
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

	jsonOutputBytes, err := toJson(headerKeys, rows)

	if err != nil {
		t.Error(err)
	}

	if string(jsonOutputBytes) != expected {
		t.Errorf("given %s, want %s", jsonOutputBytes, expected)
	}
}

func TestConvert(t *testing.T) {

	t.Run("should convert with standard delimiter and with header", func(t *testing.T) {

		csvInput := `
first_name,last_name
eder,costa
teste,teste
`

		jsonExpected := fmt.Sprintln(`[{"first_name":"eder","last_name":"costa"},{"first_name":"teste","last_name":"teste"}]`)

		tmpFile, err := ioutil.TempFile("", "*")

		if err != nil {
			t.Fatal(err)
		}

		defer os.Remove(tmpFile.Name())

		err = convert(strings.NewReader(csvInput), tmpFile, ",", false)

		if err != nil {
			t.Error(err)
		}

		jsonOutputBytes, err := ioutil.ReadFile(tmpFile.Name())

		if err != nil {
			t.Fatal(err)
		}

		if jsonExpected != string(jsonOutputBytes) {
			t.Errorf("given %s, want %s", string(jsonOutputBytes), jsonExpected)
		}
	})
}
