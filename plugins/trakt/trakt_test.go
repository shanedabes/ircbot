package trackt

import "testing"

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
	got := j.Latest()
	expected := ee.String()

	if got != expected {
		t.Errorf("got %q, want %q", got, expected)
	}
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
			got := tc.e.String()

			if got != tc.expected {
				t.Errorf("got %q, want %q", got, tc.expected)
			}
		})
	}
}

func TestEpisode(t *testing.T) {
	got := e.String()
	expected := "01x02 - test ep"

	if got != expected {
		t.Errorf("got %q, want %q", got, expected)
	}
}

func TestShow(t *testing.T) {
	got := s.String()
	expected := "test show"

	if got != expected {
		t.Errorf("got %q, want %q", got, expected)
	}
}

func TestMovie(t *testing.T) {
	got := m.String()
	expected := "test movie (2020)"

	if got != expected {
		t.Errorf("got %q, want %q", got, expected)
	}
}
