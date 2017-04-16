package modes

import (
	"os"
    "stued/terminal"
)

func CommandMode(term *terminal.Window, r rune) (terminal.InputHandler, error) {
	var next terminal.InputHandler
	cmdbuff := ""
	next = func(term *terminal.Window, r rune) (terminal.InputHandler, error) {
		switch r {
		case '\n':
			return process(term, cmdbuff)
		case terminal.CTRL_C:
			term.SetStatus("")
			return EditMode, nil
		default:
			cmdbuff += string([]rune{r})
		}
		term.SetStatus("cmd: " + cmdbuff)
		return next, nil
	}
	return next(term, r)
}

func process(term *terminal.Window, cmdbuff string) (terminal.InputHandler, error) {
	switch cmdbuff {
	case "q":
		if term.Editor.Dirty {
			term.SetStatus("Use q! to quit without saving")
			return EditMode, nil
		}
		return nil, nil
	case "q!":
		return nil, nil
	case "w":
		f, err := os.Create(term.Name)
		if err != nil {
			return EditMode, err
		}
		defer f.Close()
		err = term.Editor.WriteTo(f)
		if err == nil {
			term.Editor.Dirty = false
			term.SetStatus("file saved")
		} else {
			term.SetStatus("error saving: " + err.Error())
		}
		return EditMode, err
	default:
		term.SetStatus("unknown command")
	}
	return EditMode, nil
}
