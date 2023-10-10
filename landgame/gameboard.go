package landgame

import (
	"fmt"

	"github.com/fatih/color"
)

const boardsize int = 12

type Gameboard struct {
	board   [boardsize][boardsize]Block
	players []Player
}

// Intended for initial configuration of gameboard
func NewGameboard() Gameboard {
	var gb Gameboard
	// for i := 0; i < boardsize; i++ {
	// 	for j := 0; j < boardsize; j++ {
	// 	}
	// }
	return gb
}

func (gb Gameboard) String() string {
	bsh := ""
	for i, p := range gb.players {
		bsh += fmt.Sprintf("Player %d: %v\n", i, p)
	}
	bs := ""
	for y := 0; y < boardsize; y++ {
		for x := 0; x < boardsize; x++ {
			// bs += fmt.Sprintf("[%+v](%2d,%2d)", gb.board[y][x], x, y)
			bs += fmt.Sprintf("[%+v]", gb.board[y][x])
		}
		bs += "\n"
	}
	return bsh + "\n" + bs
}

type Block struct {
	Marker     string
	Belongs_to *Player
}

func (b Block) String() string {
	if b.Belongs_to != nil {
		return ColorString(fmt.Sprintf("%1s", b.Marker), b.Belongs_to.Color)
	}
	return fmt.Sprintf("%1s", b.Marker)
}

var (
	OpenBlock   Block = Block{}
	RobBlock    Block = Block{Marker: "R"}
	AttackBlock Block = Block{Marker: "A"}
	HomeBlock   Block = Block{Marker: "H"}
)

type Color int

const (
	Black Color = iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

func (c Color) String() string {
	return [...]string{"Black", "Red", "Green", "Yellow", "Blue", "Magenta", "Cyan", "White"}[c]
}

func ColorString(s string, c Color) string {
	// log.Default().Printf("print '%s' in %s", s, c)
	switch c {
	case Red:
		return color.RedString(s)
	case Green:
		return color.GreenString(s)
	case Yellow:
		return color.YellowString(s)
	case Blue:
		return color.BlueString(s)
	case Magenta:
		return color.MagentaString(s)
	case Cyan:
		return color.CyanString(s)
	case White:
		return color.WhiteString(s)
	}
	return s
}

type Card int

const (
	SwapCard Card = iota
	RobCard
	AttackCard
)

func (c Card) String() string {
	return [...]string{"Swap", "Rob", "Attack"}[c]
}

type Player struct {
	Name        string
	Color       Color
	SwapCards   []Card
	RobCards    []Card
	AttackCards []Card
}

func (p Player) String() string {
	return ColorString(fmt.Sprintf("%10s | Swap Cards:%d, Rob Cards:%d, Attack Cards:%d", p.Name, len(p.SwapCards), len(p.RobCards), len(p.AttackCards)), p.Color)
}
