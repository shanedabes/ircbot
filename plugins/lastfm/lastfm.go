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

var api_key = os.Getenv("IRC_LASTFM_API")

type lastfmJson struct {
	Recenttracks Recenttracks `json:"recenttracks"`
}

func (j lastfmJson) String() string {
	return j.Recenttracks.String()
}

type Recenttracks struct {
	User   User    `json:"@attr"`
	Tracks []Track `json:"track"`
}

func (r Recenttracks) String() string {
	track := r.Tracks[0]

	return fmt.Sprintf(" %s %s %s ", r.User, track.Action(), track)
}

type User struct {
	User string `json:"user"`
}

func (u User) String() string {
	return u.User
}

type Artist struct {
	Name string `json:"#text"`
}

func (a Artist) String() string {
	return a.Name
}

type TrackAttr struct {
	Nowplaying string `json:"nowplaying"`
}

type Album struct {
	Name string `json:"#text"`
}

func (a Album) String() string {
	return a.Name
}

type Track struct {
	Artist     Artist     `json:"artist"`
	Nowplaying *TrackAttr `json:"@attr,omitempty"`
	Album      Album      `json:"album"`
	Name       string     `json:"name"`
}

func (t Track) String() string {
	return fmt.Sprintf("%s - %s (%s)", t.Artist, t.Name, t.Album)
}

func (t Track) Action() string {
	if t.Nowplaying == nil {
		return "last listened to"
	} else {
		return "is listening to"
	}
}

func lastfm(command *bot.Cmd) (msg string, err error) {
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
