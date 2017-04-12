package main

import (
	"errors"
)

func EditMode(term *Terminal, r rune) (InputHandler, error) {
	term.Status = ""
	switch r {
	case ESC:
		term.Status = "cmd:"
		return CommandMode, nil
	case BACKSPACE, CTRL_H:
		term.Editor.RemoveRuneBack()
	case ARROW_UP:
		term.Editor.MoveUp(1)
	case ARROW_DOWN:
		term.Editor.MoveDown(1)
	case ARROW_LEFT:
		term.Editor.MoveLeft()
	case ARROW_RIGHT:
		term.Editor.MoveRight()
	case PAGE_UP:
		term.Editor.MoveUp(term.Rows)
	case PAGE_DOWN:
		term.Editor.MoveDown(term.Rows)
	case CTRL_C: // fixme
		return nil, errors.New("CTRL_C")
	default:
		term.Editor.AddRune(r)
	}
	return EditMode, nil
}
