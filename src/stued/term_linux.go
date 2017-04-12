// +build linux

package main

// #cgo CFLAGS: -D_XOPEN_SOURCE_EXTENDED
// #cgo LDFLAGS: -lncursesw -ltinfo
// #include <stdlib.h>
// #include <wchar.h>
// #include <ncursesw/ncurses.h>
// #include <locale.h>
import "C"

import (
	"unicode/utf8"
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
	empty := C.CString("")
	C.setlocale(C.LC_ALL, empty)
	C.free(unsafe.Pointer(empty))
	C.initscr()
	C.raw()
	C.keypad(C.stdscr, true)
	C.noecho()
}

func endRaw() {
	C.endwin()
}

func getrune() rune {
	var c C.wint_t
	C.wget_wch(C.stdscr, &c) // needs _XOPEN_SOURCE_EXTENDED
	return rune(c)
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

func addrune(r rune) {
	var buff [5]byte
	if utf8.EncodeRune(buff[:], r) > 4 {
		panic("rune too long?")
	}
	C.addstr((*C.char)(unsafe.Pointer(&buff[0])))
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
