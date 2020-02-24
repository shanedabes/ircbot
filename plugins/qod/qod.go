package qod

import (
	"fmt"

	"github.com/go-chat-bot/bot"
	"github.com/go-chat-bot/plugins/web"
)

const (
	quotesURL = "http://quotes.rest/qod.json"
)

// JSON represents the return from quotes.rest
type JSON struct {
	Contents Contents `json:"contents"`
}

func (j JSON) String() string {
	return j.Contents.String()
}

// Contents represents the list of quotes
type Contents struct {
	Quotes []Quote `json:"quotes"`
}

func (c Contents) String() string {
	if len(c.Quotes) < 1 {
		return "No quotes found"
	}
	return c.Quotes[0].String()
}

// Quote represents a single quote
type Quote struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

func (q Quote) String() string {
	return fmt.Sprintf("%s\n- %s", q.Quote, q.Author)
}

func qod(command *bot.Cmd) (msg string, err error) {
	j := &JSON{}
	err = web.GetJSON(quotesURL, j)

	if err != nil {
		return "", err
	}

	return j.String(), nil
}

func init() {
	bot.RegisterCommand(
		"qod",
		"Post today's quote from quotes.rest",
		"",
		qod,
	)
}
