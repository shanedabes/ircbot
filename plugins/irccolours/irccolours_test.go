package irccolours

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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
			assert.Equal(t, tc.f.String(), tc.expected)
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
		assert.Equal(t, len(got), 14, "Expected 14 lines")
	})

	t.Run("Colours looped", func(t *testing.T) {
		for i := 0; i < 13; i++ {
			expected := Colour(fmt.Sprintf("%02d", i+2))
			assert.Equal(t, got[i].Fg, expected)
		}
		assert.Equal(t, got[13].Fg, Colour("02"))
	})
}

func TestFormattedTextToStringList(t *testing.T) {
	fs := []FormattedText{
		FormattedText{Text: "test"},
		FormattedText{Text: "test2"},
	}
	got := FormattedTextToStringList(fs)

	for n, f := range fs {
		assert.Equal(t, f.String(), got[n])
	}
}
