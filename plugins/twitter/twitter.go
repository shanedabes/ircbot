package twitter

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-chat-bot/bot"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

var (
	consumerKey    = os.Getenv("IRC_TWITTER_CONSUMER_KEY")
	consumerSecret = os.Getenv("IRC_TWITTER_CONSUMER_SECRET")
	accessToken    = os.Getenv("IRC_TWITTER_ACCESS_TOKEN")
	accessSecret   = os.Getenv("IRC_TWITTER_ACCESS_SECRET")
)

func twit(command *bot.Cmd) (msg string, err error) {
	arg := strings.Split(command.RawArgs, " ")[0]

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	user, _, err := client.Users.Show(&twitter.UserShowParams{
		ScreenName: arg,
	})

	if err != nil {
		return "", err
	}

	if user.StatusesCount == 0 {
		return fmt.Sprintf("User %s has no tweets", arg), nil
	}

	t := strings.Split(user.Status.Text, "\n")[0]

	return t, nil
}

func init() {
	bot.RegisterCommand(
		"twitter",
		"Post user's last tweet",
		"sharktamer",
		twit,
	)
}
