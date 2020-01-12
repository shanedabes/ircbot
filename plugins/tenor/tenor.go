package tenor

import (
	"fmt"
	"math/rand"
	"net/url"

	"github.com/go-chat-bot/bot"
	"github.com/go-chat-bot/plugins/web"
)

const (
	tenor_api_url_base = "https://api.tenor.com/v1/search"
)

func search_url(term string, api_key string) (url string) {
	temp := "%s?q=%s&key=%s&limit=10"

	return fmt.Sprintf(temp, tenor_api_url_base, term, api_key)
}

type tenor_json struct {
	Results []struct {
		Media []struct {
			Gif struct {
				URL string `json:"url"`
			} `json:"gif"`
		} `json:"media"`
		Title string `json:"title"`
	} `json:"results"`
}

func tenor(command *bot.Cmd) (msg string, err error) {
	api_key := "UTX9QGBLPC26"
	data := &tenor_json{}
	msg = url.QueryEscape(command.RawArgs)

	err = web.GetJSON(search_url(msg, api_key), data)

	if err != nil {
		return "", err
	}

	if len(data.Results) == 0 {
		return "No gifs found.", nil
	}

	index := rand.Intn(len(data.Results))
	return fmt.Sprintf(data.Results[index].Media[0].Gif.URL), nil
}

func init() {
	bot.RegisterCommand(
		"gif",
		"Post a random gif from tenor",
		"louie spence",
		tenor,
	)
}
