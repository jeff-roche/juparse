package lgr

import (
	"fmt"
	"io"

	"github.com/TwiN/go-color"
)

const (
	FAILURE string = "Failed"
	WARNING string = "Warning"
	SKIPPED string = "Skipped"
	PASSED  string = "Passed"
)

func LogTestStatus(level, msg string, wrt io.Writer, useColor bool) {
	var output string
	if useColor {
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

		output = fmt.Sprintf("[%s] %s", color.Ize(c, level), msg)
	} else {
		output = fmt.Sprintf("[%s] %s", level, msg)
	}

	fmt.Fprintln(wrt, output)
}
