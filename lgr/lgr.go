package lgr

import (
	"fmt"

	"github.com/TwiN/go-color"
)

const (
	FAILURE string = "Failed"
	WARNING string = "Warning"
	SKIPPED string = "Skipped"
	PASSED  string = "Passed"
)

func LogTestStatus(level, msg string) {
	c := color.White

	switch level {
	case FAILURE:
		c = color.Red
	case WARNING:
		c = color.Yellow
	case SKIPPED:
		c = color.Gray
	case PASSED:
		c = color.Green
	}

	fmt.Printf("[%s] %s\n", color.Ize(c, level), msg)
}
