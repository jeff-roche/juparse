package parser

import (
	"encoding/xml"
	"io"

	"github.com/jeff-roche/juparse/lgr"
)

type TestSuite struct {
	XMLName   xml.Name   `xml:"testsuite"`
	Name      string     `xml:"name,attr"`
	Tests     int        `xml:"tests,attr"`
	Skipped   int        `xml:"skipped,attr"`
	Failures  int        `xml:"failures,attr"`
	Time      string     `xml:"time,attr"`
	TestCases []TestCase `xml:"testcase"`
}

func (ts TestSuite) Passed() int {
	return ts.Tests - ts.Skipped - ts.Failures
}

type TestCase struct {
	XMLName     xml.Name    `xml:"testcase"`
	Name        string      `xml:"name,attr"`
	Time        float64     `xml:"time,attr"`
	Output      string      `xml:"system-out"`
	Skipped     SkippedTest `xml:"skipped"`
	Failure     string      `xml:"failure"`
	OutputColor bool
}

func (tc TestCase) Print(wrt io.Writer, color bool) {
	lvl := lgr.PASSED

	if tc.WasSkipped() {
		lvl = lgr.SKIPPED
	}

	if tc.Failed() {
		lvl = lgr.FAILURE
	}

	lgr.LogTestStatus(lvl, tc.Name, wrt, color)
}

func (tc TestCase) WasSkipped() bool {
	return tc.Skipped.Message != ""
}

func (tc TestCase) Failed() bool {
	return tc.Failure != ""
}

func (tc TestCase) Passed() bool {
	return !tc.WasSkipped() && !tc.Failed()
}

type SkippedTest struct {
	XMLName xml.Name `xml:"skipped"`
	Message string   `xml:"message,attr"`
	Value   string   `xml:",chardata"`
}

func Parse(data []byte) (*TestSuite, error) {
	suite := &TestSuite{}

	if err := xml.Unmarshal(data, suite); err != nil {
		return nil, err
	}

	return suite, nil
}
