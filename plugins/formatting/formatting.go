package formatting

import (
	"fmt"
)

func ListToLines(list []string, lineLength int) (lines []string) {
	line := ""
	for _, i := range list {
		if len(fmt.Sprintf("%s, %s", line, i)) > lineLength {
			lines = append(lines, line)
			line = ""
		}
		if line == "" {
			line = i
		} else {
			line = fmt.Sprintf("%s, %s", line, i)
		}
	}
	lines = append(lines, line)
	return
}
