package main

import (
	"fmt"
	"os"
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

	OnTerm(startRaw)
	defer OnTerm(endRaw)

	term, err := NewTerminal(os.Args[1], f)
	f.Close()
	if err != nil {
		term.Status =
			fmt.Sprintf("error reading %s: %s", os.Args[1], err.Error())
	}
	term.ProcessInput(EditMode)
	return 0
}

func main() { os.Exit(run()) }
