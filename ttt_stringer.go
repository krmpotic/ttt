package main

import "fmt"

const (
	red    = "\033[1;31m"
	green  = "\033[1;32m"
	clrRst = "\033[0m"
)

func (b board) String() (s string) {
	// show field number if empty
	f := func(i int) string {
		if b[i] == none {
			return fmt.Sprintf("%d", i)
		}
		return b[i].String()
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
		w, d, l := g.Analyze()
		s += fmt.Sprintf(" %s%v%s", green, w, clrRst)
		s += fmt.Sprintf(" %v", d)
		s += fmt.Sprintf(" %s%v%s", red, l, clrRst)
		s += "\n"
	}
	return s
}

func (p player) String() string {
	switch p {
	case x:
		return green + "X" + clrRst
	case o:
		return red + "O" + clrRst
	default:
		return " "
	}
}
