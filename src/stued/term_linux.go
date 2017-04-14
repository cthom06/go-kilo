// +build linux

package main

// #cgo CFLAGS: -D_XOPEN_SOURCE_EXTENDED
// #cgo LDFLAGS: -lncursesw -ltinfo
// #include <stdlib.h>
// #include <wchar.h>
// #include <ncursesw/ncurses.h>
// #include <locale.h>
// #include <sys/epoll.h>
import "C"

import (
	"runtime"
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
	RESIZE      = C.KEY_RESIZE

	ERR = C.ERR
)

var outChan chan (func()) = make(chan (func()), 16)
var inChan chan (func()) = make(chan (func()), 1)
var epollfd C.int

func OnTerm(f func()) {
	c := make(chan struct{}, 1)
	outChan <- func() {
		f()
		c <- struct{}{}
	}
	<-c
}

func runOnInChan(f func()) {
	c := make(chan struct{}, 1)
	inChan <- func() {
		f()
		c <- struct{}{}
	}
	<-c
}

func init() {
	go func() {
		var epollev C.struct_epoll_event
		epollfd = C.epoll_create1(0)
		if epollfd == -1 {
			panic("couldn't epoll_create")
		}
		epollev.events = C.EPOLLIN
		// work around union
		*((*C.int)(unsafe.Pointer(&epollev.data))) = 0
		if C.epoll_ctl(epollfd, C.EPOLL_CTL_ADD, 0, &epollev) == -1 {
			panic("couldn't epoll_ctl")
		}
		runtime.LockOSThread()
		for {
			select {
			case f := <-outChan:
				f()
			case f := <-inChan:
				f()
			}
		}
	}()
}

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

func getrune() (r rune) {
	var c C.wint_t
	var ev C.struct_epoll_event
	n := C.epoll_wait(epollfd, &ev, 1, -1)
	if n == -1 {
		panic("couldn't epoll_wait")
	} else if n == 0 {
		return rune(ERR)
	}
	runOnInChan(func() {
		if C.wget_wch(C.stdscr, &c) == ERR {
			r = rune(ERR)
		} else {
			r = rune(c)
		}
	})
	return
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

func clrtoeol() {
	C.clrtoeol()
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
