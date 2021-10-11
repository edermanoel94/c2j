package main

import (
	"testing"
)

func TestParse(t *testing.T) {

	t.Run("Parse", func(t *testing.T) {

		testCases := []struct {
			desc     string
			arg      argument
			args     []string
			expected error
		}{
			{"should use default argument", argument{delimiter: ","}, []string{}, nil},
			{"should use long delimiter argument", argument{delimiter: ";"}, []string{"--delimiter", ","}, nil},
			{"should use short delimiter argument", argument{delimiter: ";"}, []string{"-d", ","}, nil},
		}

		for _, tc := range testCases {

			err := parse(&tc.arg, tc.args...)

			if err != tc.expected {
				t.Errorf("want %v, given %v", tc.expected, err)
			}

			if tc.arg.delimiter != "," {
				t.Errorf("want \"%s\", given \"%s\"", ",", ";")
			}

		}
	})
}

func TestMappingHeaders(t *testing.T) {

}

func TestGenerateJson(t *testing.T) {

}
