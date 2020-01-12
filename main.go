package main

import (
	"os"

	"github.com/go-chat-bot/bot/irc"

	_ "github.com/go-chat-bot/plugins/url"
	_ "github.com/shanedabes/ircbot/plugins/tenor"

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
