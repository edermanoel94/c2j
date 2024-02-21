package main

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {

	testCases := []struct {
		name       string
		fHelp      bool
		fVersion   bool
		fNoHeader  bool
		fDelimiter string

		inputData string
	}{
		{
			"print usage",
			true,
			false,
			false,
			",",
			"",
		},
		{
			"print version",
			false,
			true,
			false,
			",",
			"",
		},
		{
			"run with csv data",
			false,
			false,
			false,
			",",
			"first_name,last_name,phone Charleen,Roche,253-330-9889 Jenica,Briat,393-963-9525 Julie,Josselsohn,898-929-2639 Maddalena,Bessom,479-862-0782 Collete,Feldklein,143-902-5122 ",
		},
	}

	for _, tc := range testCases {

		t.Run(tc.name, func(t *testing.T) {

			tempFile, err := os.CreateTemp("", "*")

			if err != nil {
				t.Fatalf("error on test: %v\n", err)
			}

			defer func() {
				tempFile.Close()
				os.Remove(tempFile.Name())
			}()

			_, err = tempFile.WriteString(tc.inputData)

			if err != nil {
				t.Fatalf("error on test: %v\n", err)
			}

			fHelp = tc.fHelp
			fVersion = tc.fVersion
			fNoHeader = tc.fNoHeader
			fDelimiter = tc.fDelimiter

			if fHelp || fVersion {

				run(nil)
				return
			}

			run(tempFile)

		})

	}
}
