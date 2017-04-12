package main

import (
	"fmt"
	"io"
)

type Terminal struct {
	Scroll       struct{ X, Y int }
	Rows, Cols   int
	Editor       *Editor
	Name, Status string
}

type InputHandler func(term *Terminal, r rune) (InputHandler, error)

func (term *Terminal) Render() {
	clear()
	edit := term.Editor
	rx, ry, curWidth := edit.Cursor.X, edit.Cursor.Y, 0
	if ry < len(edit.Rows) {
		rx, curWidth = edit.Rows[ry].IndexToVisible(rx)
		curWidth--
	}

	term.scrollTo(ry, rx, curWidth)

	for y, editY := 0, term.Scroll.Y; y < term.Rows-2 && editY < len(edit.Rows); y, editY = y+1, editY+1 {

		move(y, 0)
		editRow := edit.Rows[editY]

		displayed := 0
		min := term.Scroll.X
		max := min + term.Cols

		for _, r := range editRow {
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

	term.renderStatus()
	move(ry-term.Scroll.Y, rx-term.Scroll.X)
	refresh()
}

func (term *Terminal) scrollTo(ry, rx, curWidth int) {
	usedRows := term.Rows - 2
	if ry-term.Scroll.Y >= usedRows {
		term.Scroll.Y = ry - (usedRows - 1)
	} else if term.Scroll.Y > ry {
		term.Scroll.Y = ry
	}

	if rx+curWidth-term.Scroll.X >= term.Cols {
		term.Scroll.X = rx + curWidth - (term.Cols - 1)
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

func (term *Terminal) ProcessInput(mode InputHandler) (err error) {
	term.Render()
	for mode != nil && err == nil {
		mode, err = mode(term, getrune())
		term.Render()
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
