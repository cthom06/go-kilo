// +build linux

package main

// #cgo LDFLAGS: -lcurses -ltinfo
// #include <stdlib.h>
// #include <curses.h>
import "C"

import (
	"unsafe"
)

const (
	CTRL_C      = 3
	CTRL_H      = 8
	TAB         = 9
	ENTER       = 10
	RETURN      = 13
	CTRL_Q      = 17
	CTRL_S      = 19
	ESC         = 27
	BACKSPACE   = C.KEY_BACKSPACE
	ARROW_LEFT  = C.KEY_LEFT
	ARROW_RIGHT = C.KEY_RIGHT
	ARROW_UP    = C.KEY_UP
	ARROW_DOWN  = C.KEY_DOWN
	DEL_KEY     = C.KEY_DC
	HOME_KEY    = C.KEY_HOME
	END_KEY     = C.KEY_END
	PAGE_UP     = C.KEY_PPAGE
	PAGE_DOWN   = C.KEY_NPAGE
)

func startRaw() {
	C.initscr()
	C.raw()
	C.keypad(C.stdscr, true)
	C.noecho()
}

func endRaw() {
	C.endwin()
}

func getch() int {
	return int(C.getch())
}

func getWindowSize() (int, int) {
	return int(C.stdscr._maxy) + 1, int(C.stdscr._maxx) + 1
}

func move(y, x int) {
	C.move(C.int(y), C.int(x))
}

func clear() {
	C.clear()
}

func refresh() {
	C.refresh()
}

func addch(b byte) {
	C.addch(C.chtype(b))
}

func addstr(s string) {
	p := C.CString(s)
	C.addstr(p)
	C.free(unsafe.Pointer(p))
}

func startReverse() {
	C.attron(C.A_REVERSE)
}

func endReverse() {
	C.attroff(C.A_REVERSE)
}
