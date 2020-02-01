package trackt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	e = Episode{
		Season: 1,
		Number: 2,
		Title:  "test ep",
	}

	s = Show{
		Title: "test show",
	}

	m = Movie{
		Title: "test movie",
		Year:  2020,
	}

	ee = Entry{
		Type:    "episode",
		Episode: e,
		Show:    s,
	}

	em = Entry{
		Type:  "movie",
		Movie: m,
	}

	j = traktJSON{ee, em}
)

func TestJson(t *testing.T) {
	assert.Equal(t, j.Latest(), ee.String())
}

func TestEntry(t *testing.T) {
	cases := []struct {
		name     string
		e        Entry
		expected string
	}{
		{
			name:     "episode entry",
			e:        ee,
			expected: "test show 01x02 - test ep",
		},
		{
			name:     "movie entry",
			e:        em,
			expected: "test movie (2020)",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.e.String(), tc.expected)
		})
	}
}

func TestEpisode(t *testing.T) {
	expected := "01x02 - test ep"
	assert.Equal(t, e.String(), expected)
}

func TestShow(t *testing.T) {
	expected := "test show"
	assert.Equal(t, s.String(), expected)
}

func TestMovie(t *testing.T) {
	expected := "test movie (2020)"
	assert.Equal(t, m.String(), expected)
}
