package irccolours

import (
	"fmt"
)

// Colour is used to construct a coloured string for irc
type Colour string

const (
	// White formatted irc text string
	White Colour = "00"
	// Black formatted irc text string
	Black Colour = "01"
	// Blue formatted irc text string
	Blue Colour = "02"
	// Green formatted irc text string
	Green Colour = "03"
	// Red formatted irc text string
	Red Colour = "04"
	// Brown formatted irc text string
	Brown Colour = "05"
	// Magenta formatted irc text string
	Magenta Colour = "06"
	// Orange formatted irc text string
	Orange Colour = "07"
	// Yellow formatted irc text string
	Yellow Colour = "08"
	// LGreen (light green) formatted irc text string
	LGreen Colour = "09"
	// Cyan formatted irc text string
	Cyan Colour = "10"
	// LCyan (light cyan) formatted irc text string
	LCyan Colour = "11"
	// LBlue (light blue) formatted irc text string
	LBlue Colour = "12"
	// Pink formatted irc text string
	Pink Colour = "13"
	// Grey formatted irc text string
	Grey Colour = "14"
	// Gray formatted irc text string
	Gray Colour = "14"
	// LGrey (light grey) formatted irc text string
	LGrey Colour = "15"
	// LGray (light gray) formatted irc text string
	LGray Colour = "15"
	// None (default) formatted irc text string
	None Colour = "99"
)

// FormattedText represents a coloured irc string
type FormattedText struct {
	Text   string
	Fg, Bg Colour
}

func (f FormattedText) String() string {
	if f.Fg == "" {
		f.Fg = None
	}

	if f.Bg == "" {
		f.Bg = None
	}

	return fmt.Sprintf("\x03%s,%s%s\x0399,99", f.Fg, f.Bg, f.Text)
}

// ColouriseList applies colours to all strings in a list, cycling through all available
func ColouriseList(strs []string) (out []FormattedText) {
	for n, s := range strs {
		c := Colour(fmt.Sprintf("%02d", n%13+2))
		out = append(out, FormattedText{Text: s, Fg: c})
	}
	return out
}

// FormattedTextToStringList formats a list to a newline separated string
func FormattedTextToStringList(fs []FormattedText) (out []string) {
	for _, f := range fs {
		out = append(out, f.String())
	}
	return
}
