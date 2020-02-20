package clock

import (
	"fmt"
	"time"

	"github.com/go-chat-bot/bot"
)

func clock(command *bot.Cmd) (msg string, err error) {
	tz := command.RawArgs

	now := time.Now()

	loc, err := time.LoadLocation(tz)

	if err != nil {
		return "", err
	}

	t := now.In(loc).Format("Jan 1 15:04:05")

	return fmt.Sprintf("%s", t), nil
}

func init() {
	bot.RegisterCommand(
		"time",
		"Post current time",
		"Europe/London",
		clock,
	)
}
