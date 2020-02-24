package jod

import (
	"github.com/go-chat-bot/bot"
	"github.com/go-chat-bot/plugins/web"
)

const (
	url = "https://api.jokes.one/jod"
)

type json struct {
	Contents contents `json:"contents"`
}

func (j json) String() string {
	return j.Contents.String()
}

type contents struct {
	Jokes []jokes `json:"jokes"`
}

func (c contents) String() string {
	return c.Jokes[0].String()
}

type jokes struct {
	Joke joke `json:"joke"`
}

func (js jokes) String() string {
	return js.Joke.String()
}

type joke struct {
	Text string `json:"text"`
}

func (j joke) String() string {
	return j.Text
}

func jod(command *bot.Cmd) (msg string, err error) {
	j := &json{}
	err = web.GetJSON(url, j)

	if err != nil {
		return "", err
	}

	return j.String(), nil
}

func init() {
	bot.RegisterCommand(
		"jod",
		"Post the joke of the day from api.jokes.one",
		"",
		jod,
	)
}
