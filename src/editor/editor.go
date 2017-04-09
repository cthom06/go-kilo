package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

// This file is mostly ported from antirez' kilo
// as such it is subject to the license terms in
// KILO_LICENSE

const (
	TAB_WIDTH           = 4
	STATUS_DISPLAY_SECS = 5
	QUIT_TIMES          = 1
)

type erow struct {
	idx   int
	chars []byte
}

type editorConfig struct {
	cx, cy                 int
	rowoff, coloff         int
	screenrows, screencols int
	rows                   []erow
	dirty                  bool
	filename               string
	statusmsg              string
	statusmsg_time         time.Time
}

func (E *editorConfig) insertRow(at int, s []byte) {
	if at > len(E.rows) || at < 0 {
		panic("bad insertRow?")
	}
	newrows := E.rows
	if cap(newrows) == len(newrows) {
		newrows = make([]erow, len(newrows)+1, len(newrows)+4)
		copy(newrows, E.rows[0:at])
	} else {
		newrows = newrows[0 : len(newrows)+1]
	}
	copy(newrows[at+1:], E.rows[at:])
	for i := at + 1; i < len(newrows); i++ {
		newrows[i].idx++
	}
	newrows[at].chars = s
	newrows[at].idx = at
	E.rows = newrows
	E.dirty = true
}

func (E *editorConfig) delRow(at int) {
	if at < 0 || at >= len(E.rows) {
		return
	}
	for i := at + 1; i < len(E.rows); i++ {
		E.rows[i-1] = E.rows[i]
		E.rows[i-1].idx--
	}
	E.rows = E.rows[:len(E.rows)-1]
	E.dirty = true
}

func (E *editorConfig) insertChar(rowInd, colInd int, c byte) {
	row := &E.rows[rowInd]
	if colInd > len(row.chars) {
		panic("bad insertChar?")
	}
	row.chars = append(row.chars, 0)
	copy(row.chars[colInd+1:], row.chars[colInd:])
	row.chars[colInd] = c
	E.dirty = true
}

func (E *editorConfig) delChar(rowInd, colInd int) {
	row := &E.rows[rowInd]
	copy(row.chars[colInd:], row.chars[colInd+1:])
	row.chars = row.chars[:len(row.chars)-1]
	E.dirty = true
}

func (E *editorConfig) insertCharAtCursor(c byte) {
	for E.cy >= len(E.rows) {
		E.insertRow(len(E.rows), []byte{})
	}
	E.insertChar(E.cy, E.cx, c) // E is dirty
	E.cx++
}

func (E *editorConfig) insertNewlineAtCursor() {
	if E.cx == 0 || E.cy == len(E.rows) {
		E.insertRow(E.cy, []byte{}) // E is dirty
	} else {
		E.insertRow(E.cy+1, []byte{}) // E is dirty
		E.appendBytes(E.cy+1, E.rows[E.cy].chars[E.cx:])
		E.rows[E.cy].chars = E.rows[E.cy].chars[:E.cx]
	}
	E.cy++
	E.cx = 0
}

func (E *editorConfig) delCharAtCursor() {
	if E.cx == 0 && E.cy == 0 {
		return
	}
	if E.cx == 0 {
		E.cx = len(E.rows[E.cy-1].chars)
		E.appendBytes(E.cy-1, E.rows[E.cy].chars) // E is dirty
		E.delRow(E.cy)
		E.cy--
	} else {
		E.delChar(E.cy, E.cx-1) // E is dirty
		E.cx--
	}
}

func (E *editorConfig) appendBytes(rowInd int, b []byte) {
	E.rows[rowInd].chars = append(E.rows[rowInd].chars, b...)
	E.dirty = true
}

func (E *editorConfig) cursorRenderPosition() (int, int) {
	if E.cy >= len(E.rows) {
		return 0, 0
	}
	row := &E.rows[E.cy]
	cx := 0
	for i := 0; i < E.cx; i++ {
		if row.chars[i] == '\t' {
			cx = cx + TAB_WIDTH - cx%TAB_WIDTH
		} else {
			cx++
		}
	}
	return E.cy, cx
}

func (E *editorConfig) renderPositionToCursor(y, x int) (int, int) {
	if y < 0 || y >= len(E.rows) {
		return 0, 0
	}
	row := &E.rows[y]
	cx := 0
	for i := 0; i < len(row.chars); i++ {
		if row.chars[i] == '\t' {
			cx = cx + TAB_WIDTH - cx%TAB_WIDTH
		} else {
			cx++
		}
		if cx > x {
			return y, i
		}
	}
	return y, len(row.chars)
}

func (E *editorConfig) open(filename string) error {
	E.filename = filename

	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	bf := bufio.NewReader(f)
	line, err := bf.ReadBytes('\n')
	for err == nil && len(line) > 0 {
		E.insertRow(len(E.rows), line[:len(line)-1])
		line, err = bf.ReadBytes('\n')
	}
	E.dirty = false
	if err == io.EOF {
		err = nil
	}
	return err
}

func (E *editorConfig) save() error {
	f, err := os.Create(E.filename)
	if err != nil {
		return err
	}
	defer f.Close()

	bf := bufio.NewWriter(f)

	for _, v := range E.rows {
		if _, err := bf.Write(v.chars); err != nil {
			return err
		}
		if err := bf.WriteByte('\n'); err != nil {
			return err
		}
	}
	E.dirty = false
	if err := bf.Flush(); err != nil {
		return err
	}
	E.setStatusMessage("File saved")
	return nil
}

func (E *editorConfig) refreshScreen() {
	clear()
	y := 0
	ry, rx := E.cursorRenderPosition()
	for ry-E.rowoff >= E.screenrows {
		E.rowoff++
	}
	for ry-E.rowoff < 0 {
		E.rowoff--
	}
	for rx-E.coloff >= E.screencols {
		E.coloff++
	}
	for rx-E.coloff < 0 {
		E.coloff--
	}
	for ; y < E.screenrows; y++ {
		move(y, 0)
		fr := E.rowoff + y

		// ignore anything past the end of the file
		if fr < len(E.rows) {
			row := &E.rows[fr]
			l := 0
			minl := E.coloff
			maxl := minl + E.screencols
			for j := 0; j < len(row.chars) && l < maxl; j++ {
				b := row.chars[j]
				if b == TAB {
					spaces := TAB_WIDTH - l%TAB_WIDTH
					for i := 0; i < spaces; i++ {
						if l >= minl && l < maxl {
							addch(' ')
						}
						l++
					}
				} else {
					if l >= minl {
						addch(b)
					}
					l++
				}
			}
		}
	}
	move(y, 0)
	startReverse()
	d := ""
	if E.dirty {
		d = "(dirty)"
	}
	status :=
		fmt.Sprintf("%d,%d/%d - %s %s",
			rx+1,
			ry+1,
			len(E.rows),
			E.filename,
			d)
	if len(status) > E.screencols {
		status = status[:E.screencols]
	}
	addstr(status)
	for i := len(status); i < E.screencols; i++ {
		addch(' ')
	}
	endReverse()

	move(y+1, 0)
	if time.Since(E.statusmsg_time) < STATUS_DISPLAY_SECS*time.Second {
		addstr(E.statusmsg)
	}

	move(ry-E.rowoff, rx-E.coloff)
	refresh()
}

func (E *editorConfig) setStatusMessage(msg string) {
	E.statusmsg = msg
	E.statusmsg_time = time.Now()
}

func (E *editorConfig) cursorToBounds() {
	if E.cy >= len(E.rows) {
		E.cy = len(E.rows) - 1
	}
	if E.cy > -1 {
		if E.cx > len(E.rows[E.cy].chars) {
			E.cx = len(E.rows[E.cy].chars)
		}
	} else {
		E.cy = 0
		E.cx = 0
	}
}

func (E *editorConfig) moveCursor(key int) {
	var row *erow = nil
	if E.cy < len(E.rows) {
		row = &E.rows[E.cy]
	}
	switch key {
	case ARROW_LEFT:
		if E.cx == 0 {
			if E.cy > 0 {
				E.cy--
				E.cx = len(E.rows[E.cy].chars)
			}
		} else {
			E.cx--
		}
	case ARROW_RIGHT:
		if row != nil {
			if E.cx == len(row.chars) {
				if E.cy < len(E.rows)-1 {
					E.cy++
					E.cx = 0
				}
			} else {
				E.cx++
			}
		}
	case ARROW_UP, ARROW_DOWN:
		_, rx := E.cursorRenderPosition()
		if key == ARROW_UP {
			E.cy--
		} else {
			E.cy++
		}
		E.cy, E.cx = E.renderPositionToCursor(E.cy, rx)
	default:
		panic("bad key to moveCursor")
	}
	E.cursorToBounds()
}

var qtimes = QUIT_TIMES

func (E *editorConfig) processKeypress() (quit bool) {
	c := getch()
	switch c {
	case ENTER, RETURN:
		E.insertNewlineAtCursor()
	case CTRL_Q:
		if E.dirty && qtimes > 0 {
			status := fmt.Sprintf("!!! Unsaved changes: Press Ctrl-Q %d more time to discard and quit.", qtimes)
			E.setStatusMessage(status)
			qtimes--
			return false
		}
		return true
	case CTRL_S:
		E.save()
	case BACKSPACE, CTRL_H, DEL_KEY:
		E.delCharAtCursor()
	case PAGE_UP:
		E.cy = E.cy - E.screenrows
		E.cursorToBounds()
	case PAGE_DOWN:
		E.cy = E.cy + E.screenrows
		E.cursorToBounds()
	case ARROW_UP, ARROW_DOWN, ARROW_LEFT, ARROW_RIGHT:
		E.moveCursor(c)
	default:
		E.insertCharAtCursor(byte(c))
	}
	qtimes = QUIT_TIMES
	return false
}

func initEditor() *editorConfig {
	r := new(editorConfig)
	r.rows = make([]erow, 0, 128)
	r.screenrows, r.screencols = getWindowSize()
	r.screenrows = r.screenrows - 2 // issue if negative?
	return r
}

func run() int {
	args := os.Args
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <filename>\n", args[0])
		return 1
	}
	startRaw()
	defer endRaw()
	E := initEditor()
	if err := E.open(args[1]); err != nil {
		E.setStatusMessage("couldn't open " + args[1] + ": " + err.Error())
	}
	for true {
		E.refreshScreen()
		if E.processKeypress() {
			// got a quit signal
			return 0
		}
	}
	panic("unreachable")
}

func main() {
	os.Exit(run())
}
