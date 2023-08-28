package main

type player int

const (
	o    = player(-1)
	none = player(0)
	x    = player(1)
)

func (p player) other() player {
	return -p
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
