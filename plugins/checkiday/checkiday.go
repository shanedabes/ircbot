package checkiday

import (
	"github.com/go-chat-bot/bot"
	"github.com/go-chat-bot/plugins/web"

	"github.com/shanedabes/ircbot/plugins/formatting"
	"github.com/shanedabes/ircbot/plugins/irccolours"
)

const (
	checkiday_api_url = "https://checkiday.com/api/3/?d"
)

type daysJson struct {
	Holidays []struct {
		Name string `json:"name"`
	} `json:"holidays"`
}

func (dj daysJson) days() (out []string) {
	for _, d := range dj.Holidays {
		out = append(out, d.Name)
	}
	return
}

func checkiday(command *bot.Cmd) (msg string, err error) {
	data := &daysJson{}
	err = web.GetJSON(checkiday_api_url, data)

	if err != nil {
		return "", err
	}

	cl := irccolours.ColouriseList(data.days())
	dl := irccolours.FormattedTextToStringList(cl)
	return formatting.ListToLines(dl, 400), nil
}

func init() {
	bot.RegisterCommand(
		"days",
		"A daily listing of today's holidays",
		"",
		checkiday,
	)
}
