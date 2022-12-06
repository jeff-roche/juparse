package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jeff-roche/juparse/parser"
)

var (
	file    *string = flag.String("file", "", "[Required] the file to parse")
	skipped *bool   = flag.Bool("skipped", false, "Filters to show skipped tests")
	passed  *bool   = flag.Bool("passed", false, "Filters to show passed tests")
	failed  *bool   = flag.Bool("failed", false, "Filteres to show failed tests")
	verbose *bool   = flag.Bool("v", false, "Print the reason it failed or was skipped")
)

func main() {
	flag.Parse()

	if *file == "" {
		fmt.Println("No file specified. Please specify a file to parse")
		os.Exit(1)
	}

	suite := parseTestResults()

	processOutput(suite)
}

func parseTestResults() *parser.TestSuite {
	// Read the file
	xmlFile, err := os.Open(*file)
	if err != nil {
		fmt.Printf("Unable to read the file (%s): %s", *file, err)
		os.Exit(1)
	}

	defer xmlFile.Close()

	data, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Printf("Error reading the file contents: %s", err)
		os.Exit(1)
	}

	suite, err := parser.Parse(data)
	if err != nil {
		fmt.Printf("Unable to parse the xml: %s", err)
		os.Exit(1)
	}

	return suite
}

func processOutput(suite *parser.TestSuite) {
	showAll := !*skipped && !*passed && !*failed

	fmt.Printf("Processing test suite '%s'\n", suite.Name)
	fmt.Printf("Total Tests: %d\tPassed: %d\tSkipped: %d\tFailed: %d\n", suite.Tests, suite.Passed(), suite.Skipped, suite.Failures)

	for _, tc := range suite.TestCases {
		if showAll {
			tc.Print()
		}

		if *skipped && tc.WasSkipped() {
			tc.Print()

			if *verbose {
				fmt.Printf("\t%s\n", tc.Skipped.Message)
			}
		}

		if *failed && tc.Failed() {
			tc.Print()

			if *verbose {
				fmt.Printf("\t%s\n", tc.Failure)
			}
		}

		if *passed && tc.Passed() {
			tc.Print()
		}
	}
}
