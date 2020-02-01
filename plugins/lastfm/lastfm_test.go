package lastfm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	u = User{
		User: "testuser",
	}

	ar = Artist{
		Name: "testartist",
	}

	al = Album{
		Name: "testalbum",
	}

	tr = Track{
		Artist: ar,
		Album:  al,
		Name:   "testtrack",
	}

	trp = Track{
		Artist: ar,
		Album:  al,
		Name:   "testtrack",
		Nowplaying: &TrackAttr{
			Nowplaying: "blah",
		},
	}

	rt = Recenttracks{
		User:   u,
		Tracks: []Track{tr},
	}

	rtn = Recenttracks{
		Tracks: []Track{},
	}
)

func TestString(t *testing.T) {
	cases := []struct {
		name     string
		obj      fmt.Stringer
		expected string
	}{
		{
			name:     "User",
			obj:      u,
			expected: "testuser",
		},
		{
			name:     "Artist",
			obj:      ar,
			expected: "testartist",
		},
		{
			name:     "Album",
			obj:      al,
			expected: "testalbum",
		},
		{
			name:     "Track",
			obj:      tr,
			expected: "testartist - testtrack (testalbum)",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.obj.String(), tc.expected)
		})
	}
}

func TestAction(t *testing.T) {
	cases := []struct {
		name     string
		t        Track
		expected string
	}{
		{
			name:     "Playing track",
			t:        trp,
			expected: "is listening to",
		},
		{
			name:     "Not playing track",
			t:        tr,
			expected: "last listened to",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.t.action()

			assert.Equal(t, got, tc.expected)
		})
	}
}

func TestRecentTracks(t *testing.T) {
	cases := []struct {
		name     string
		r        Recenttracks
		expected string
	}{
		{
			name:     "Valid result",
			r:        rt,
			expected: " testuser last listened to testartist - testtrack (testalbum) ",
		},
		{
			name:     "Invalid results",
			r:        rtn,
			expected: "No tracks found for user",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.r.String()

			assert.Equal(t, got, tc.expected)
		})
	}
}
