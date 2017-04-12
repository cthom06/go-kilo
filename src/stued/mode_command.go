package main

import (
	"os"
)

func CommandMode(term *Terminal, r rune) (InputHandler, error) {
	var next InputHandler
	cmdbuff := ""
	next = func(term *Terminal, r rune) (InputHandler, error) {
		switch r {
		case '\n':
			return process(term, cmdbuff)
		case CTRL_C:
			term.Status = ""
			return EditMode, nil
		default:
			cmdbuff += string([]rune{r})
		}
		term.Status = "cmd: " + cmdbuff
		return next, nil
	}
	return next(term, r)
}

func process(term *Terminal, cmdbuff string) (InputHandler, error) {
	switch cmdbuff {
	case "q":
		if term.Editor.Dirty {
			term.Status = "Use q! to quit without saving"
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
			term.Status = "file saved"
		} else {
			term.Status = "error saving: " + err.Error()
		}
		return EditMode, err
	default:
		term.Status = "unknown command"
	}
	return EditMode, nil
}
