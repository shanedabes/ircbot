package checkiday

import (
	"github.com/go-chat-bot/bot"
	"github.com/go-chat-bot/plugins/web"
	"github.com/shanedabes/ircbot/plugins/formatting"
	"github.com/shanedabes/ircbot/plugins/irccolours"
)

const (
	checkidayAPIURL = "https://checkiday.com/api/3/?d"
)

type daysJSON struct {
	Days []Day `json:"holidays"`
}

func (ds daysJSON) List() (out []string) {
	for _, d := range ds.Days {
		out = append(out, d.Name)
	}
	return
}

// Day represents a single day result from checkiday
type Day struct {
	Name string `json:"name"`
}

func (d Day) String() string {
	return d.Name
}

func checkiday(command *bot.Cmd) (msg string, err error) {
	data := &daysJSON{}
	err = web.GetJSON(checkidayAPIURL, data)

	if err != nil {
		return "", err
	}

	cl := irccolours.ColouriseList(data.List())
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
