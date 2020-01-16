package choose

import (
	"math/rand"
	"strings"
	"time"

	"github.com/go-chat-bot/bot"
)

func choose(command *bot.Cmd) (msg string, err error) {
	args := command.RawArgs
	words := strings.Split(args, "/")

	rand.Seed(time.Now().Unix())
	index := rand.Intn(len(words))
	word := words[index]

	return word, nil
}

func init() {
	bot.RegisterCommand(
		"choose",
		"Choose one from options",
		"red lolly/yellow lolly/green lolly",
		choose,
	)
}
