package main

import "fmt"

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
