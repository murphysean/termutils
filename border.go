package termutils

import (
	//"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

//go:generate stringer -type=Border
type Border int

const (
	//BorderLight Light Border
	BorderLight Border = iota
	//BorderLightArc Light Arc Border
	BorderLightArc
	//BorderLightDoubleLine Light Border with double line
	BorderLightDoubleLine
	//BorderLightDoubleDash Light Border with double dash
	BorderLightDoubleDash
	//BorderLightTripleDash Light Border with triple dash
	BorderLightTripleDash
	//BorderLightQuadDash Light Border with quad dash
	BorderLightQuadDash
	//BorderThick Thick Border
	BorderThick
	//BorderThickDoubleDash Thick border with a double dash
	BorderThickDoubleDash
	//BorderThickTripleDash Thick border with a triple dash
	BorderThickTripleDash
	//BorderThickQuadDash Thick border with a quad dash
	BorderThickQuadDash
)

// DrawBoxBorder will draw a box border
func DrawBoxBorder(x, y, w, h int, t Border) {
	drawBoxBorder(x, y, w, h, t, termbox.ColorDefault, termbox.ColorDefault)
}

// DrawBoxBorderColor will draw a box border with color attributes
func DrawBoxBorderColor(x, y, w, h int, t Border, fg, bg termbox.Attribute) {
	drawBoxBorder(x, y, w, h, t, fg, bg)
}

func drawBoxBorder(x, y, w, h int, t Border, fg, bg termbox.Attribute) {
	var hr, vr, tl, tr, bl, br rune
	switch t {
	case BorderLight:
		tl = '┌'
		tr = '┐'
		bl = '└'
		br = '┘'
		hr = '─'
		vr = '│'
	case BorderLightArc:
		tl = '╭'
		tr = '╮'
		bl = '╰'
		br = '╯'
		hr = '─'
		vr = '│'
	case BorderLightDoubleLine:
		tl = '╔'
		tr = '╗'
		bl = '╚'
		br = '╝'
		hr = '═'
		vr = '║'
	case BorderLightDoubleDash:
		tl = '┌'
		tr = '┐'
		bl = '└'
		br = '┘'
		hr = '╌'
		vr = '╎'
	case BorderLightTripleDash:
		tl = '┌'
		tr = '┐'
		bl = '└'
		br = '┘'
		hr = '┄'
		vr = '┆'
	case BorderLightQuadDash:
		tl = '┌'
		tr = '┐'
		bl = '└'
		br = '┘'
		hr = '┈'
		vr = '┊'
	case BorderThick:
		tl = '┏'
		tr = '┓'
		bl = '┗'
		br = '┛'
		hr = '━'
		vr = '┃'
	case BorderThickDoubleDash:
		tl = '┏'
		tr = '┓'
		bl = '┗'
		br = '┛'
		hr = '╍'
		vr = '╏'
	case BorderThickTripleDash:
		tl = '┏'
		tr = '┓'
		bl = '┗'
		br = '┛'
		hr = '┅'
		vr = '┇'
	case BorderThickQuadDash:
		tl = '┏'
		tr = '┓'
		bl = '┗'
		br = '┛'
		hr = '┉'
		vr = '┋'
	}
	for i := x; i < x+w; i++ {
		termbox.SetCell(i, y, hr, fg, bg)
		termbox.SetCell(i, y+h-1, hr, fg, bg)
	}
	for i := y; i < y+h; i++ {
		termbox.SetCell(x, i, vr, fg, bg)
		termbox.SetCell(x+w, i, vr, fg, bg)
	}
	termbox.SetCell(x, y, tl, fg, bg)
	termbox.SetCell(x+w, y, tr, fg, bg)
	termbox.SetCell(x, y+h-1, bl, fg, bg)
	termbox.SetCell(x+w, y+h-1, br, fg, bg)
}
