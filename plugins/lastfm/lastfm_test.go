package lastfm

import (
	"github.com/stretchr/testify/assert"
	"testing"
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

func TestUser(t *testing.T) {
	got := u.String()
	expected := "testuser"

	assert.Equal(t, got, expected)
}

func TestArtist(t *testing.T) {
	got := ar.String()
	expected := "testartist"

	assert.Equal(t, got, expected)
}

func TestAlbum(t *testing.T) {
	got := al.String()
	expected := "testalbum"

	assert.Equal(t, got, expected)
}

func TestTrack(t *testing.T) {
	got := tr.String()
	expected := "testartist - testtrack (testalbum)"

	assert.Equal(t, got, expected)
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
			got := tc.t.Action()

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
