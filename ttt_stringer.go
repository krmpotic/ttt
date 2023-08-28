package main

import "fmt"

func (b board) String() (s string) {
	// show field number if empty
	f := func(i int) string {
		if b[i] == none {
			return fmt.Sprintf("%d", i)
		}
		return b[i].String()
	}
	if clearScreen {
		s += clrScr
	}
	s += fmt.Sprintf(" %s ║ %s ║ %s\n", f(0), f(1), f(2))
	s += fmt.Sprintf("═══╬═══╬═══\n")
	s += fmt.Sprintf(" %s ║ %s ║ %s\n", f(3), f(4), f(5))
	s += fmt.Sprintf("═══╬═══╬═══\n")
	s += fmt.Sprintf(" %s ║ %s ║ %s\n", f(6), f(7), f(8))
	return s
}

func (g *game) String() (s string) {
	s = fmt.Sprintf("\n%s", g.board)
	if showAnalysis {
		w, d, l := g.analyze(-1)
		s += fmt.Sprintf(" %s%v%s", green, w, clrRst)
		s += fmt.Sprintf(" %v", d)
		s += fmt.Sprintf(" %s%v%s", red, l, clrRst)
		s += "\n"
	}
	return s
}
