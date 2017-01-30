package termutils

import (
	"github.com/nsf/termbox-go"
)

// DrawScrollbar This will draw a thin scrollbar
func DrawScrollbar(x, y, h, min, max, winmin, winmax int) {
	if min > max || min > winmin || min > winmax {
		return
	}
	if max < min || max < winmax || max < winmin {
		return
	}

	//ltr := '╷'
	//ttr := '╻'
	lmr := '│'
	tmr := '┃'
	tatr := '╿'
	tabr := '╽'
	//lbr := '╷'
	//tbr := '╹'

	//We get 2 'pixels' per char block, with a loss of 1 block -- .5 at top and .5 at bottom
	sectionSize := float64(max-min) / (float64(h) * 2)
	//TODO An offset when things don't min at 0

	for i, c := y, float64(0); i < y+h; i++ {
		top := true
		bottom := true
		//Block Start
		bs := sectionSize * c
		//Block End
		be := sectionSize * (c + 1)

		if (float64(winmin) < bs && float64(winmax) < bs) || (float64(winmin) > be && float64(winmax) > be) {
			//Draw thin
			top = false
		}

		bs = be
		be = sectionSize * (c + 2)

		if (float64(winmin) < bs && float64(winmax) < bs) || (float64(winmin) > be && float64(winmax) > be) {
			bottom = false
		}

		if top && bottom {
			termbox.SetCell(x, i, tmr, termbox.ColorDefault, termbox.ColorDefault)
		} else if top {
			termbox.SetCell(x, i, tatr, termbox.ColorDefault, termbox.ColorDefault)
		} else if bottom {
			termbox.SetCell(x, i, tabr, termbox.ColorDefault, termbox.ColorDefault)
		} else {
			termbox.SetCell(x, i, lmr, termbox.ColorDefault, termbox.ColorDefault)
		}

		c += 2
	}
}
