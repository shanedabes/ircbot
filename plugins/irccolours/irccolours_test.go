package irccolours

import (
	"fmt"
	"testing"
)

func TestColour(t *testing.T) {
	cases := []struct {
		name     string
		f        FormattedText
		expected string
	}{
		{
			name:     "Foreground",
			f:        FormattedText{Text: "t", Fg: Blue},
			expected: "\x0302,99t\x0399,99",
		},
		{
			name:     "Foreground and background",
			f:        FormattedText{Text: "t", Fg: Red, Bg: Black},
			expected: "\x0304,01t\x0399,99",
		},
		{
			name:     "Background",
			f:        FormattedText{Text: "t", Bg: Green},
			expected: "\x0399,03t\x0399,99",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.f.String()

			if got != tc.expected {
				t.Errorf("got %q, want %q", got, tc.expected)
			}
		})
	}
}

func TestColouriseList(t *testing.T) {
	var ss []string
	for i := 0; i < 14; i++ {
		ss = append(ss, "t")
	}

	got := ColouriseList(ss)

	t.Run("14 elements", func(t *testing.T) {
		if len(got) != 14 {
			t.Errorf("expected 14 elements")
		}
	})

	t.Run("Colours looped", func(t *testing.T) {
		for i := 0; i < 13; i++ {
			if g := got[i]; g.Fg != Colour(fmt.Sprintf("%02d", i+2)) {
				t.Errorf("got %s, expected %02d", g.Fg, i+2)
			}
		}

		if got[13].Fg != Colour("02") {
			t.Errorf("got %s, expected 02", got[13].Fg)
		}
	})
}
