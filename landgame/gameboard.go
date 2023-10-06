package landgame

import (
	"fmt"

	"github.com/fatih/color"
)

const boardsize int = 12

type Gameboard [boardsize][boardsize]Block

func NewGameboard() Gameboard {
	var gb Gameboard
	for i := 0; i < boardsize; i++ {
		for j := 0; j < boardsize; j++ {
		}
	}
	return gb
}

func (gb Gameboard) String() string {
	bsh := ""
	bs := ""
	for y := 0; y < boardsize; y++ {
		for x := 0; x < boardsize; x++ {
			bs += fmt.Sprintf("[%+v](%2d,%2d)", gb[y][x], x, y)
		}
		bs += "\n"
	}
	return bsh + "\n" + bs
}

type Block struct {
	Marker     string
	Belongs_to Player
}

func (b Block) String() string {
	return b.Belongs_to.ColorString(fmt.Sprintf("%3s", b.Marker))
}

var (
	OpenBlock   Block = Block{Marker: "", Belongs_to: Player{}}
	RobBlock    Block = Block{Marker: " R ", Belongs_to: Player{}}
	AttackBlock Block = Block{Marker: " A ", Belongs_to: Player{}}
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

func (p Player) ColorString(s string) string {
	switch p.Color {
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
	return p.ColorString(p.Name)
}
