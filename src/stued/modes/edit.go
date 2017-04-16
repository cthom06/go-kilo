package modes

import (
	"errors"
    "stued/terminal"
)

func EditMode(term *terminal.Window, r rune) (terminal.InputHandler, error) {
	// term.Status = ""
	switch r {
	case terminal.ESC:
		term.SetStatus("cmd:")
		return CommandMode, nil
	case terminal.BACKSPACE, terminal.CTRL_H:
		term.Editor.RemoveRuneBack()
	case terminal.ARROW_UP:
		term.Editor.MoveUp(1)
	case terminal.ARROW_DOWN:
		term.Editor.MoveDown(1)
	case terminal.ARROW_LEFT:
		term.Editor.MoveLeft()
	case terminal.ARROW_RIGHT:
		term.Editor.MoveRight()
	case terminal.PAGE_UP:
		term.Editor.MoveUp(term.Rows)
	case terminal.PAGE_DOWN:
		term.Editor.MoveDown(term.Rows)
	case terminal.CTRL_C: // fixme
		return nil, errors.New("CTRL_C")
	default:
		term.Editor.AddRune(r)
	}
	return EditMode, nil
}
