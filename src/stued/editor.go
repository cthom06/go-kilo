package main

import (
	"bufio"
	"io"
)

const (
	TAB_WIDTH = 4
)

type Row []rune

type Editor struct {
	Cursor struct{ X, Y int }
	Rows   []Row
	Dirty  bool
	Events chan EditorEvent
}

func (edit *Editor) InsertRow(at int, row Row) {
	edit.Rows = append(edit.Rows, nil)
	copy(edit.Rows[at+1:], edit.Rows[at:])
	edit.Rows[at] = row
	edit.Dirty = true
}

func (edit *Editor) RemoveRow(at int) {
	copy(edit.Rows[at:], edit.Rows[at+1:])
	edit.Rows = edit.Rows[:len(edit.Rows)-1]
	edit.Dirty = true
}

func (row *Row) Insert(at int, r rune) {
	*row = append(*row, 0)
	copy((*row)[at+1:], (*row)[at:])
	(*row)[at] = r
}

func (row *Row) Remove(at int) {
	copy((*row)[at:], (*row)[at+1:])
	(*row) = (*row)[:len(*row)-1]
}

func (row Row) IndexToVisible(ind int) (vis, width int) {
	sum := 0
	for i := 0; i < ind && i < len(row); i++ {
		r := row[i]
		if r == '\t' {
			sum = sum + TAB_WIDTH - sum%TAB_WIDTH
		} else {
			w, _ := runeWidth(r)
			sum = sum + w
		}
	}
	if ind >= len(row) {
		return sum, 1
	} else if row[ind] == '\t' {
		return sum, TAB_WIDTH - sum%TAB_WIDTH
	}
	w, _ := runeWidth(row[ind])
	return sum, w
}

func (row Row) VisibleToIndex(vis int) int {
	sum := 0
	for i := 0; i < len(row); i++ {
		if row[i] == '\t' {
			sum = sum + TAB_WIDTH - sum%TAB_WIDTH
		} else {
			w, _ := runeWidth(row[i])
			sum = sum + w
		}

		if sum > vis {
			return i
		}
	}
	return len(row)
}

func (edit *Editor) AddRune(r rune) {
	if r == '\n' {
		edit.addNewline()
	} else {
		y := edit.Cursor.Y
		for y >= len(edit.Rows) {
			edit.InsertRow(len(edit.Rows), nil)
		}
		edit.Rows[y].Insert(edit.Cursor.X, r)
		edit.Cursor.X++
		edit.Dirty = true
	}
	edit.Events <- RuneInserted{Cursor: edit.Cursor, Rune: r}
}

func (edit *Editor) RemoveRuneBack() {
	cursor := edit.Cursor
	x, y := cursor.X, cursor.Y
	var removed rune
	if x == 0 {
		if y == 0 {
			return // nothing to backspace
		}
		edit.Cursor.X = len(edit.Rows[y-1])
		edit.Cursor.Y--
		if y < len(edit.Rows) {
			edit.Rows[y-1] = append(edit.Rows[y-1], edit.Rows[y]...)
			edit.RemoveRow(y) // dirty
		}
		removed = '\n'
	} else {
		removed = edit.Rows[y][x-1]
		edit.Rows[y].Remove(x - 1)
		edit.Cursor.X--
		edit.Dirty = true
	}
	edit.Events <- RuneRemoved{Cursor: cursor, Rune: removed}
}

func (edit *Editor) addNewline() {
	x, y := edit.Cursor.X, edit.Cursor.Y
	if x == 0 || y == len(edit.Rows) {
		edit.InsertRow(y, nil) // dirty
	} else {
		edit.InsertRow(y+1, append([]rune{}, edit.Rows[y][x:]...)) // dirty
		edit.Rows[y] = edit.Rows[y][:x]
	}
	edit.Cursor.X = 0
	edit.Cursor.Y++
}

func NewEditor(content io.Reader) (ret Editor, err error) {
	br, ok := content.(*bufio.Reader)
	if !ok {
		br = bufio.NewReader(content)
	}
	runeBuff := make([]rune, 0, 64)
	var r rune
	for r, _, err = br.ReadRune(); err == nil; r, _, err = br.ReadRune() {
		if r == '\n' {
			// copy out runeBuff
			ret.Rows = append(ret.Rows, append([]rune{}, runeBuff...))
			runeBuff = runeBuff[:0]
		} else {
			runeBuff = append(runeBuff, r)
		}
	}
	if len(runeBuff) != 0 {
		ret.Rows = append(ret.Rows, append([]rune{}, runeBuff...))
	}
	if err == io.EOF {
		err = nil
	}
	ret.Events = make(chan EditorEvent, 16)
	return
}

func (edit *Editor) WriteTo(w io.Writer) error {
	bw, ok := w.(*bufio.Writer)
	if !ok {
		bw = bufio.NewWriter(w)
	}
	for _, row := range edit.Rows {
		for _, r := range row {
			if _, err := bw.WriteRune(r); err != nil {
				return err
			}
		}
		if _, err := bw.WriteRune('\n'); err != nil {
			return err
		}
	}
	return bw.Flush()
}

func (edit *Editor) MoveDown(n int) {
	x, y := edit.Cursor.X, edit.Cursor.Y
	if n < 0 || y+n >= len(edit.Rows) {
		n = len(edit.Rows) - y - 1
	}
	if y < len(edit.Rows)-1 {
		x, _ = edit.Rows[y].IndexToVisible(x)
		if y+n < len(edit.Rows) {
			edit.Cursor.X = edit.Rows[y+n].VisibleToIndex(x)
		} else {
			edit.Cursor.X = 0
		}
		edit.Cursor.Y += n
	}
	edit.Events <- CursorMoved{X: edit.Cursor.X - x, Y: edit.Cursor.Y - y}
}

func (edit *Editor) MoveUp(n int) {
	x, y := edit.Cursor.X, edit.Cursor.Y
	if n < 0 || y-n < 0 {
		n = y
	}
	if y > 0 {
		if y < len(edit.Rows) {
			x, _ = edit.Rows[y].IndexToVisible(x)
			edit.Cursor.X = edit.Rows[y-n].VisibleToIndex(x)
		} else {
			edit.Cursor.X = 0
		}
		edit.Cursor.Y -= n
	}
	edit.Events <- CursorMoved{X: edit.Cursor.X - x, Y: edit.Cursor.Y - y}
}

// MoveLeft and MoveRight will try to skip zero-width characters

func (edit *Editor) MoveLeft() {
	x, y := edit.Cursor.X, edit.Cursor.Y
	if x == 0 {
		if y > 0 {
			edit.Cursor.X = len(edit.Rows[y-1])
			edit.Cursor.Y--
		}
	} else {
		edit.Cursor.X--
		if w, _ := runeWidth(edit.Rows[y][x-1]); w == 0 {
			edit.MoveLeft()
		}
	}
	edit.Events <- CursorMoved{X: edit.Cursor.X - x, Y: edit.Cursor.Y - y}
}

func (edit *Editor) MoveRight() {
	x, y := edit.Cursor.X, edit.Cursor.Y
	if y < len(edit.Rows) {
		if x == len(edit.Rows[y]) {
			if y+1 < len(edit.Rows) {
				edit.Cursor.X = 0
				edit.Cursor.Y++
				if len(edit.Rows[y+1]) > 0 {
					if w, _ := runeWidth(edit.Rows[y+1][0]); w == 0 {
						edit.MoveRight()
					}
				}
			}
		} else {
			edit.Cursor.X++
			if x+1 < len(edit.Rows[y]) {
				if w, _ := runeWidth(edit.Rows[y][x+1]); w == 0 {
					edit.MoveRight()
				}
			}
		}
	}
	edit.Events <- CursorMoved{X: edit.Cursor.X - x, Y: edit.Cursor.Y - y}
}
