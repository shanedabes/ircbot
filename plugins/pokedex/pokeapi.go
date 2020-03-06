package pokedex

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-chat-bot/bot"
	"github.com/go-chat-bot/plugins/web"
)

const (
	example       = "slowbro"
	apiURL        = "https://pokeapi.co/api/v2/"
	speciesAPIURL = "pokemon-species/%s"
)

var (
	errIncorrectUsage = errors.New("species name must be included")
)

func pk(command *bot.Cmd) (msg string, err error) {
	w := strings.Split(command.RawArgs, " ")
	if len(w) < 2 {
		return "", errIncorrectUsage
	}
	species, game := w[0], w[1]

	url := apiURL + fmt.Sprintf(speciesAPIURL, species)

	j := &Species{}
	err = web.GetJSON(url, j)
	if err != nil {
		return "", err
	}

	f, err := j.FlavorTextEntries.Select("en", game)
	if err != nil {
		return "", err
	}

	return f.FlavorText, nil
}

func init() {
	bot.RegisterCommand(
		"pokedex",
		"Return data on selected pokemon",
		example,
		pk,
	)
}
