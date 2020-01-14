package lastfm

import (
	"fmt"
	"net/url"
	"os"

	"github.com/go-chat-bot/bot"
	"github.com/go-chat-bot/plugins/web"
)

const (
	lastfmRecentTracksApiURL = "https://ws.audioscrobbler.com/2.0" +
		"?method=user.getrecenttracks" +
		"&user=%s&api_key=%s&format=json&limit=1"
)

type lastfmJson struct {
	Recenttracks struct {
		Attr struct {
			User string `json:"user"`
		} `json:"@attr"`
		Track []struct {
			Album struct {
				Text string `json:"#text"`
			} `json:"album"`
			Attr struct {
				NowPlaying string `json:"nowplaying"`
			} `json:"@attr"`
			Artist struct {
				Text string `json:"#text"`
			} `json:"artist"`
			Name string `json:"name"`
		} `json:"track"`
	} `json:"recenttracks"`
}

func (l lastfmJson) String() string {
	user := l.Recenttracks.Attr.User
	track := l.Recenttracks.Track[0]
	album := track.Album.Text
	artist := track.Artist.Text
	trackName := track.Name

	action := "last listened to"
	if track.Attr.NowPlaying == "true" {
		action = "is listening to"
	}

	return fmt.Sprintf(
		"♫  %s %s %s - %s (%s) ♫",
		user, action, artist, trackName, album,
	)
}

func lastfm(command *bot.Cmd) (msg string, err error) {
	api_key := os.Getenv("IRC_LASTFM_API")
	msg = url.QueryEscape(command.RawArgs)
	url := fmt.Sprintf(lastfmRecentTracksApiURL, msg, api_key)

	data := &lastfmJson{}
	err = web.GetJSON(url, data)

	if err != nil {
		return "", err
	}

	return data.String(), nil
}

func init() {
	bot.RegisterCommand(
		"lastfm",
		"Post user's last played song on last.fm",
		"sharktamer",
		lastfm,
	)
}
