package lastfm

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestLastfmJsonString(t *testing.T) {
	jf := `{
		  "recenttracks": {
			"@attr": {
			  "user": "user"
			},
			"track": [
			  {
				"artist": {
				  "#text": "artist"
				},
				"@attr": {
				  "nowplaying": "%s"
				},
				"album": {
				  "#text": "album"
				},
				"name": "track"
			  }
			]
		  }
		}`

	cases := []struct {
		j        string
		expected string
	}{
		{
			fmt.Sprintf(jf, "true"),
			"♫ user is listening to artist - track (album) ♫",
		},
		{
			fmt.Sprintf(jf, ""),
			"♫ user last listened to artist - track (album) ♫",
		},
	}

	for _, tc := range cases {
		data := &lastfmJson{}
		_ = json.Unmarshal([]byte(tc.j), &data)

		got := data.String()
		if got != tc.expected {
			t.Errorf("got %q, want %q", got, tc.expected)
		}
	}
}
