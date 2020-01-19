package lastfm

import (
	"fmt"
	"net/url"
	"os"

	"github.com/go-chat-bot/bot"
	"github.com/go-chat-bot/plugins/web"
)

const (
	lastfmRecentTracksAPIURL = "https://ws.audioscrobbler.com/2.0" +
		"?method=user.getrecenttracks" +
		"&user=%s&api_key=%s&format=json&limit=1"
)

var apiKey = os.Getenv("IRC_LASTFM_API")

type lastfmJSON struct {
	Recenttracks Recenttracks `json:"recenttracks"`
}

func (j lastfmJSON) String() string {
	return j.Recenttracks.String()
}

// Recenttracks represents the metadata returned in the lastfm json
type Recenttracks struct {
	User   User    `json:"@attr"`
	Tracks []Track `json:"track"`
}

func (r Recenttracks) String() string {
	if len(r.Tracks) == 0 {
		return "No tracks found for user"
	}

	track := r.Tracks[0]

	return fmt.Sprintf(" %s %s %s ", r.User, track.action(), track)
}

// User represents the user information returned in the lastfm json
type User struct {
	User string `json:"user"`
}

func (u User) String() string {
	return u.User
}

// Artist represents the artist information returned in the lastfm json
type Artist struct {
	Name string `json:"#text"`
}

func (a Artist) String() string {
	return a.Name
}

// TrackAttr contains the track metadata, with now playing information
type TrackAttr struct {
	Nowplaying string `json:"nowplaying"`
}

// Album represents the album information returned in the lastfm json
type Album struct {
	Name string `json:"#text"`
}

func (a Album) String() string {
	return a.Name
}

// Track represents the track information returned in the lastfm json
type Track struct {
	Artist     Artist     `json:"artist"`
	Nowplaying *TrackAttr `json:"@attr,omitempty"`
	Album      Album      `json:"album"`
	Name       string     `json:"name"`
}

func (t Track) String() string {
	return fmt.Sprintf("%s - %s (%s)", t.Artist, t.Name, t.Album)
}

func (t Track) action() string {
	if t.Nowplaying == nil {
		return "last listened to"
	}
	return "is listening to"
}

func lastfm(command *bot.Cmd) (msg string, err error) {
	msg = url.QueryEscape(command.RawArgs)
	url := fmt.Sprintf(lastfmRecentTracksAPIURL, msg, apiKey)

	j := &lastfmJSON{}
	err = web.GetJSON(url, j)

	if err != nil {
		return "", err
	}

	return j.String(), nil
}

func init() {
	bot.RegisterCommand(
		"lastfm",
		"Post user's last played song on last.fm",
		"sharktamer",
		lastfm,
	)
}
