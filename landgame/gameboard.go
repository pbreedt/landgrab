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
	for _, p := range gb.players {
		bsh += ColorString(fmt.Sprintf("Player %s in %s\n", p.Name, p.Color), p.Color)
	}
	bs := ""
	for y := 0; y < boardsize; y++ {
		for x := 0; x < boardsize; x++ {
			bs += fmt.Sprintf("[%+v](%2d,%2d)", gb.board[y][x], x, y)
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
		return ColorString(fmt.Sprintf("%3s", b.Marker), b.Belongs_to.Color)
	}
	return fmt.Sprintf("%3s", b.Marker)
}

var (
	OpenBlock   Block = Block{Marker: ""}
	RobBlock    Block = Block{Marker: " R "}
	AttackBlock Block = Block{Marker: " A "}
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

func (e Color) String() string {
	return [...]string{"Red", "Green", "Yellow", "Blue", "Magenta", "Cyan", "White"}[e]
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

type Player struct {
	Name  string
	Color Color
}

func (p Player) String() string {
	return ColorString(p.Name, p.Color)
}
