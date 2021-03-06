package formatting

import (
	"fmt"
)

// ListToLines converts a list of strings to a newline separated string, without breaking entries across lines
func ListToLines(list []string, lineLength int) (lines string) {
	line := ""
	for _, i := range list {
		if len(fmt.Sprintf("%s, %s", line, i)) > lineLength {
			if lines == "" {
				lines = line
			} else {
				lines = fmt.Sprintf("%s\n%s", lines, line)
			}
			line = ""
		}

		if line == "" {
			line = i
		} else {
			line = fmt.Sprintf("%s, %s", line, i)
		}
	}
	lines = fmt.Sprintf("%s\n%s", lines, line)
	return
}
