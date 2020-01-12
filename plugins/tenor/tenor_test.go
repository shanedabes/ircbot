package tenor

import (
	"fmt"
	"testing"
)

func Test_search_url(t *testing.T) {
	got := search_url("a", "b")
	want := fmt.Sprintf("%s?q=a&key=b&limit=10", tenor_api_url_base)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
