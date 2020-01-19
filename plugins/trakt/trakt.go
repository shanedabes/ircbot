package trackt

import (
	"fmt"
	"net/url"
	"os"

	"github.com/go-chat-bot/bot"
	"github.com/shanedabes/ircbot/plugins/web"
)

const (
	traktAPIURL = "https://api.trakt.tv/users/%s/history"
)

var apiKey = os.Getenv("IRC_TRAKT_API")

type traktJSON []Entry

func (tj traktJSON) Latest() string {
	return tj[0].String()
}

// Entry represents a single entry, film or episode, in the trakt json
type Entry struct {
	Type    string  `json:"type"`
	Episode Episode `json:"episode,omitempty"`
	Show    Show    `json:"show,omitempty"`
	Movie   Movie   `json:"movie,omitempty"`
}

func (e Entry) String() string {
	if e.Type == "episode" {
		return fmt.Sprintf("%s %s", e.Show, e.Episode)
	}

	if e.Type == "movie" {
		return e.Movie.String()
	}

	return "unknown"
}

// Episode represents a single episode in the trakt json
type Episode struct {
	Season int    `json:"season"`
	Number int    `json:"number"`
	Title  string `json:"title"`
}

func (e Episode) String() string {
	return fmt.Sprintf("%02dx%02d - %s", e.Season, e.Number, e.Title)
}

// Show contains the parent show information for the returned episode
type Show struct {
	Title string `json:"title"`
}

func (s Show) String() string {
	return s.Title
}

// Movie represents a single movie in the trakt json
type Movie struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
}

func (m Movie) String() string {
	return fmt.Sprintf("%s (%d)", m.Title, m.Year)
}

func trakt(command *bot.Cmd) (msg string, err error) {
	args := url.QueryEscape(command.RawArgs)
	url := fmt.Sprintf(traktAPIURL, args)

	j := &traktJSON{}
	headers := map[string]string{
		"Content-Type":      "application/json",
		"trakt-api-version": "2",
		"trakt-api-key":     apiKey,
	}
	err = web.GetJSONWithHeaders(url, headers, j)

	if err != nil {
		return "", err
	}

	out := fmt.Sprintf("%s last watched: %s", args, j.Latest())

	return out, nil
}

func init() {
	bot.RegisterCommand(
		"trakt",
		"Post user's last played film or episode on trakt",
		"sharktamer",
		trakt,
	)
}
