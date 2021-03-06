package main

import (
	"os"

	"github.com/go-chat-bot/bot/irc"

	_ "github.com/go-chat-bot/plugins/url"
	_ "github.com/shanedabes/ircbot/plugins/checkiday"
	_ "github.com/shanedabes/ircbot/plugins/choose"
	_ "github.com/shanedabes/ircbot/plugins/clock"
	_ "github.com/shanedabes/ircbot/plugins/donger"
	_ "github.com/shanedabes/ircbot/plugins/fact"
	_ "github.com/shanedabes/ircbot/plugins/jod"
	_ "github.com/shanedabes/ircbot/plugins/lastfm"
	_ "github.com/shanedabes/ircbot/plugins/pokedex"
	_ "github.com/shanedabes/ircbot/plugins/qod"
	_ "github.com/shanedabes/ircbot/plugins/qwantz"
	_ "github.com/shanedabes/ircbot/plugins/talk"
	_ "github.com/shanedabes/ircbot/plugins/tenor"
	_ "github.com/shanedabes/ircbot/plugins/trakt"
	_ "github.com/shanedabes/ircbot/plugins/twitter"
	_ "github.com/shanedabes/ircbot/plugins/wikipedia"
	_ "github.com/shanedabes/ircbot/plugins/wordnik"

	"strings"
)

func main() {
	irc.Run(&irc.Config{
		Server:        os.Getenv("IRC_SERVER"),
		Channels:      strings.Split(os.Getenv("IRC_CHANNELS"), ","),
		User:          os.Getenv("IRC_USER"),
		Nick:          os.Getenv("IRC_NICK"),
		Password:      os.Getenv("IRC_PASSWORD"),
		UseTLS:        false,
		TLSServerName: os.Getenv("IRC_TLS_SERVER_NAME"),
		Debug:         os.Getenv("DEBUG") != ""})
}
