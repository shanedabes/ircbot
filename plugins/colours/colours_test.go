package colours

import (
	"testing"
)

func TestColour(t *testing.T) {
	cases := []struct {
		f        FormattedText
		expected string
	}{
		{FormattedText{Text: "t", Fg: Blue}, "\x0302,99t\x0399,99"},
		{FormattedText{Text: "t", Fg: Red, Bg: Black}, "\x0304,01t\x0399,99"},
		{FormattedText{Text: "t", Bg: Green}, "\x0399,03t\x0399,99"},
	}

	for _, tc := range cases {
		got := tc.f.String()

		if got != tc.expected {
			t.Errorf("got %q, want %q", got, tc.expected)
		}
	}
}
