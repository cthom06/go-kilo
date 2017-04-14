package main

import (
	"fmt"
	"io"
	"time"
)

type Terminal struct {
	Scroll       struct{ X, Y int }
	Rows, Cols   int
	Editor       *Editor
	Name, Status string
}

type InputHandler func(term *Terminal, r rune) (InputHandler, error)

func (term *Terminal) ProcessEditorEvent(ev EditorEvent) {
	switch ev := ev.(type) {
	case CursorMoved:
		scroll, rx, _ := term.needToScroll()
		if scroll {
			term.Redraw()
		} else {
			term.renderStatus()
			move(term.Editor.Cursor.Y-term.Scroll.Y, rx-term.Scroll.X)
			refresh()
		}
	case RuneInserted:
		scroll, rx, _ := term.needToScroll()
		if scroll || ev.Rune == '\n' {
			term.Redraw()
		} else {
			term.redrawLine(term.Editor.Cursor.Y)
			term.renderStatus()
			move(term.Editor.Cursor.Y-term.Scroll.Y, rx)
			refresh()
		}
	case RuneRemoved:
		scroll, rx, _ := term.needToScroll()
		if scroll || ev.Rune == '\n' {
			term.Redraw()
		} else {
			term.redrawLine(term.Editor.Cursor.Y)
			term.renderStatus()
			move(term.Editor.Cursor.Y-term.Scroll.Y, rx)
			refresh()
		}
	default:
		term.Redraw()
	}
}

func (term *Terminal) redrawLine(y int) {
	ry := y - term.Scroll.Y
	if ry < 0 || ry >= term.Rows-2 {
		return // offscreen
	}
	move(ry, 0)
	clrtoeol()
	term.writeLine(term.Editor.Rows[y])
}

func (term *Terminal) writeLine(row Row) {
	displayed := 0
	min := term.Scroll.X
	max := min + term.Cols

	for _, r := range row {
		if displayed >= max {
			break // no more room
		}

		if r == '\t' {
			spaces := TAB_WIDTH - displayed%TAB_WIDTH
			for i := 0; i < spaces; i++ {
				if displayed >= min && displayed < max {
					addrune(' ')
				}
				displayed++
			}
		} else {
			rsize, ashex := runeWidth(r)
			if displayed >= min && displayed+rsize <= max {
				if ashex {
					startReverse()
					addstr(fmt.Sprintf("<%X>", r))
					endReverse()
				} else {
					addrune(r)
				}
			}

			// check if character straddled a scroll boundary
			if displayed < min && displayed+rsize > min {
				n := displayed + rsize - min
				startReverse()
				for i := 0; i < n; i++ {
					addrune('<')
				}
				endReverse()
			} else if displayed <= max && displayed+rsize > max {
				n := max - displayed
				startReverse()
				for i := 0; i < n; i++ {
					addrune('>')
				}
				endReverse()
			}
			displayed += rsize
		}
	}
}

func (term *Terminal) needToScroll() (needed bool, rx, cw int) {
	edit := term.Editor
	sx, sy := term.Scroll.X, term.Scroll.Y
	ex, ey, w := edit.Cursor.X, edit.Cursor.Y, 1
	if ey < len(edit.Rows) {
		ex, w = edit.Rows[ey].IndexToVisible(ex)
	}
	needed = ex+w-1 < sx || ex >= sx+term.Cols || ey < sy || ey > sy+term.Rows-2
	return needed, ex, w
}

func (term *Terminal) Redraw() {
	clear()
	edit := term.Editor
	ry := edit.Cursor.Y
	scroll, rx, curswidth := term.needToScroll()
	curswidth--

	if scroll {
		term.scrollTo(ry, rx, curswidth)
	}
	maxy := term.Rows - 2
	maxrow := len(edit.Rows)
	for y, editY := 0, term.Scroll.Y; y < maxy && editY < maxrow; y, editY = y+1, editY+1 {

		move(y, 0)
		term.writeLine(edit.Rows[editY])
	}

	term.renderStatus()
	move(ry-term.Scroll.Y, rx-term.Scroll.X)
	refresh()
}

func (term *Terminal) scrollTo(ry, rx, curswidth int) {
	usedRows := term.Rows - 2
	if ry-term.Scroll.Y >= usedRows {
		term.Scroll.Y = ry - (usedRows - 1)
	} else if term.Scroll.Y > ry {
		term.Scroll.Y = ry
	}

	if rx+curswidth-term.Scroll.X >= term.Cols {
		term.Scroll.X = rx + curswidth - (term.Cols - 1)
	} else if term.Scroll.X > rx {
		term.Scroll.X = rx
	}
}

func (term *Terminal) renderStatus() {
	// line 1
	move(term.Rows-2, 0)
	startReverse()
	d := ""
	if term.Editor.Dirty {
		d = "(unsaved)"
	}
	status :=
		fmt.Sprintf("%d,%d/%d - %s %s",
			term.Editor.Cursor.X+1,
			term.Editor.Cursor.Y+1,
			len(term.Editor.Rows),
			term.Name,
			d)
	if len(status) > term.Cols {
		status = status[:term.Cols]
	}
	addstr(status)
	for i := len(status); i < term.Cols; i++ {
		addrune(' ')
	}
	endReverse()
	// line 2
	move(term.Rows-1, 0)
	addstr(term.Status)
}

func (term *Terminal) startInputChan() chan rune {
	c := make(chan rune, 8)
	go func() {
		for {
			r := getrune()
			if r == ERR {
				<-time.After(100 * time.Millisecond)
			} else {
				c <- r
			}
		}
	}()
	return c
}

func (term *Terminal) ProcessInput(initMode InputHandler) (err error) {
	term.Redraw()
	input := term.startInputChan()
	mode := initMode
	for mode != nil && err == nil {
		select {
		case ev := <-term.Editor.Events:
			OnTerm(func() { term.ProcessEditorEvent(ev) })
		case r := <-input:
			mode, err = mode(term, r)
		}
	}
	return
}

func NewTerminal(name string, contents io.Reader) (ret Terminal, err error) {
	ed, e := NewEditor(contents)
	if e != nil {
		err = e
		return
	}
	ret.Editor = &ed
	ret.Name = name
	ret.Rows, ret.Cols = getWindowSize()
	return
}
