package fact

import (
	"github.com/go-chat-bot/bot"
	"github.com/go-chat-bot/plugins/web"
)

const (
	url = "https://uselessfacts.jsph.pl/random.json?language=en"
)

type json struct {
	Text string `json:"text"`
}

func (j json) String() string {
	return j.Text
}

func fact(command *bot.Cmd) (msg string, err error) {
	j := &json{}
	err = web.GetJSON(url, j)

	if err != nil {
		return "", err
	}

	return j.String(), nil
}

func init() {
	bot.RegisterCommand(
		"fact",
		"Post a random fact from uselessfacts.jsph.pl",
		"",
		fact,
	)
}
