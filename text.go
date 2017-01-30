package termutils

import (
	"github.com/nsf/termbox-go"
)

// DrawText will draw a line of text
func DrawText(x, y int, text string, fg, bg termbox.Attribute) {
	i := x
	for _, r := range text {
		//fmt.Println("Setting", r, i, y)
		termbox.SetCell(i, y, r, fg, bg)
		i++
	}
}

func DrawList(x, y, width, height int, list []string, fg, bg termbox.Attribute) {

}

func DrawLine(x, y, width int, text string, fg, bg termbox.Attribute) {
	i := x
	for _, r := range text {
		if i-x > width {
			break
		}
		termbox.SetCell(i, y, r, fg, bg)
	}
}

func DrawMultiline(x, y, width, height int, text string, fg, bg termbox.Attribute) {
	i, j := x, y
	for _, r := range text {
		if r == '\n' {
			j++
			continue
		}
		if i-x > width {
			j++
		}
		if j-y > height {
			break
		}
		termbox.SetCell(i, j, r, fg, bg)
	}
}
