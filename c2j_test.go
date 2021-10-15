package main

import "testing"

func TestMappingHeaders(t *testing.T) {

	t.Run("should mapping headers", func(t *testing.T) {
		lines := [][]string{
			{"first_name", "last_name"},
		}

		headers, err := mappingHeaders(lines)

		if err != nil {
			t.Fatal(err)
		}

		if headers[0] != "first_name" {
			t.Errorf("want %s, given %s", "first_name", headers[0])
		}

		if headers[1] != "last_name" {
			t.Errorf("want %s, given %s", "last_name", headers[1])
		}
	})

	t.Run("should give a error for not have any columns but have one line", func(t *testing.T) {
		lines := [][]string{
			{},
		}

		_, err := mappingHeaders(lines)

		if err == nil {
			t.Fatal(err)
		}
	})
}

func TestGenerateJson(t *testing.T) {

	lines := [][]string{
		{"first_name", "last_name"},
		{"eder", "costa"},
		{"teste", "teste"},
	}

	expected := `[{"first_name":"eder","last_name":"costa"},{"first_name":"teste","last_name":"teste"}]`

	jsonString, err := generateJson(lines)

	if err != nil {
		t.Error(err)
	}

	if jsonString != expected {
		t.Errorf("given %s, want %s", jsonString, expected)
	}
}
