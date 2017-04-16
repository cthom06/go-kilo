package main

import (
	"fmt"
	"os"
	"stued/modes"
	"stued/terminal"
)

func run() int {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <filename>\n", os.Args[0])
		return 1
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s: couldn't open %s: %s\n",
			os.Args[0],
			os.Args[1],
			err.Error())
		return 2
	}

	terminal.OnTerm(terminal.StartRaw)
	defer terminal.OnTerm(terminal.EndRaw)

	term, err := terminal.NewWindow(os.Args[1], f)
	f.Close()
	if err != nil {
		term.SetStatus(
			fmt.Sprintf("error reading %s: %s", os.Args[1], err.Error()))
	}
	term.ProcessInput(modes.EditMode)
	return 0
}

func main() { os.Exit(run()) }
