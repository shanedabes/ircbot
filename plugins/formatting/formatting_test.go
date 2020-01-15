package formatting

import (
	"reflect"
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
	expected := []string{
		"test line",
		"another test line",
		"test, test",
		"a longer test line",
	}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %q, want %q", got, expected)
	}
}
