package main

type Player int

const (
	O    = Player(-1)
	None = Player(0)
	X    = Player(1)
)

func (p Player) Other() Player {
	return -p
}

func (p Player) String() string {
	switch p {
	case X:
		return green + "X" + clrRst
	case O:
		return red + "O" + clrRst
	default:
		return " "
	}
}
