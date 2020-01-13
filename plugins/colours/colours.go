package colours

import (
	"fmt"
)

type Colour string

const (
	White   Colour = "00"
	Black   Colour = "01"
	Blue    Colour = "02"
	Green   Colour = "03"
	Red     Colour = "04"
	Brown   Colour = "05"
	Magenta Colour = "06"
	Orange  Colour = "07"
	Yellow  Colour = "08"
	LGreen  Colour = "09"
	Cyan    Colour = "10"
	LCyan   Colour = "11"
	LBlue   Colour = "12"
	Pink    Colour = "13"
	Grey    Colour = "14"
	Gray    Colour = "14"
	LGrey   Colour = "15"
	LGray   Colour = "15"
	None    Colour = "99"
)

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
