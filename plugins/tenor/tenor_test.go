package tenor

import (
	"fmt"
	"testing"
)

func Test_search_url(t *testing.T) {
	got := searchURL("a", "b")
	want := fmt.Sprintf("%s?q=a&key=b&limit=10", tenorAPIURLbase)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
