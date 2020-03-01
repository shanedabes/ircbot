package wordnik

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-chat-bot/bot"
	"github.com/go-chat-bot/plugins/web"
)

const (
	apiRoot     = "https://api.wordnik.com/v4/"
	wodEndPoint = "words.json/wordOfTheDay?api_key=%s"
)

var (
	apiKey    = os.Getenv("IRC_WORDNIK_API")
	wodAPIURL = apiRoot + fmt.Sprintf(wodEndPoint, apiKey)
)

type json struct {
	Word        string  `json:"word"`
	Definitions []child `json:"definitions"`
	Examples    []child `json:"examples"`
	Note        string  `json:"note"`
}

func (j json) Definition() string {
	sl := []string{}
	for _, s := range j.Definitions {
		sl = append(sl, s.Text)
	}
	return strings.Join(sl, " / ")
}

func (j json) Example() string {
	sl := []string{}
	for _, s := range j.Examples {
		sl = append(sl, fmt.Sprintf(`"%s"`, s.Text))
	}
	return strings.Join(sl, "\n")
}

func (j json) String() (out string) {
	return strings.Join([]string{
		fmt.Sprintf("%s - %s", j.Word, j.Definition()),
		j.Example(),
		j.Note,
	}, "\n")
}

type child struct {
	Text string `json:"text"`
}

func wod(command *bot.Cmd) (string, error) {
	j := &json{}
	err := web.GetJSON(wodAPIURL, j)

	if err != nil {
		return "", err
	}

	return j.String(), nil
}

func init() {
	bot.RegisterCommand(
		"wod",
		"Post the word of the day from wordnik",
		"",
		wod,
	)
}
