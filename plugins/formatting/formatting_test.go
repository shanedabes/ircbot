package formatting

import (
	"strings"
	"testing"
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

	if got != expected {
		t.Errorf("got %q, want %q", got, expected)
	}
}
