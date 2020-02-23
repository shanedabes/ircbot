package qwantz

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/go-chat-bot/bot"
)

const (
	qwantzURL = "http://www.qwantz.com"
	reg       = `(http:\/\/www\.qwantz\.com\/index.php\?comic=\d+)">([^<]+)`
)

type comic struct {
	url  string
	desc string
}

func (c comic) String() string {
	return fmt.Sprintf("%s - %s", c.desc, c.url)
}

func httpError(code int, status string) error {
	return fmt.Errorf("status code error: %d %s", code, status)
}

func getPage(url string) (page string, err error) {
	res, err := http.Get("http://www.qwantz.com")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", httpError(res.StatusCode, res.Status)
	}

	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

func extractRandComic(url string) (out comic, err error) {
	page, err := getPage(url)

	if err != nil {
		return out, err
	}

	r, _ := regexp.Compile(reg)
	rr := r.FindStringSubmatch(page)
	comicURL := rr[1]
	desc := rr[2]

	return comic{
		url:  comicURL,
		desc: strings.ReplaceAll(desc, "\n", " "),
	}, nil
}

func qwantz(command *bot.Cmd) (msg string, err error) {
	comic, err := extractRandComic(qwantzURL)

	if err != nil {
		return "", err
	}

	return comic.String(), nil
}

func init() {
	bot.RegisterCommand(
		"qwantz",
		"Post random dinosaur comic",
		"",
		qwantz,
	)
}
