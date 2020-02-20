package formatting

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListToLines(t *testing.T) {
	lines := []string{
		"test line",
		"another test line",
		"test",
		"test",
		"a longer test line",
	}

	got := ListToLines(lines, 20)
	expected := strings.Join([]string{
		"test line",
		"another test line",
		"test, test",
		"a longer test line",
	}, "\n")

	assert.Equal(t, got, expected)
}
