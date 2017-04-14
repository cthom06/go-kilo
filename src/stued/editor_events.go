package main

type EditorEvent interface {
	Name() string
}

type CursorMoved struct {
	X, Y int
}

func (_ CursorMoved) Name() string {
	return "cursor moved"
}

type RuneInserted struct {
	Cursor struct{ X, Y int }
	Rune   rune
}

func (_ RuneInserted) Name() string {
	return "rune inserted"
}

type RuneRemoved struct {
	Cursor struct{ X, Y int }
	Rune   rune
}

func (_ RuneRemoved) Name() string {
	return "rune removed"
}
