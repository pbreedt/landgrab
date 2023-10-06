package landgame

import (
	"fmt"

	"github.com/fatih/color"
)

const boardsize int = 10

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
	for i := 0; i < boardsize; i++ {
		for j := 0; j < boardsize; j++ {
			bs += fmt.Sprintf("[%+v](%d,%d)", gb[i][j], i, j)
		}
		bs += "\n"
	}
	return bsh + "\n" + bs
}

type Block struct {
	Occupied_by Entity
}

func (b Block) String() string {
	return b.Occupied_by.Belongs_to.ColorString(fmt.Sprintf("%v", b.Occupied_by.Etype))
}

type EntityType int

const (
	OpenSpace EntityType = iota
	PlayerHome
	PlayerLand
	Rob
	Attack
)

func (e EntityType) String() string {
	return [...]string{"   ", " H ", " # ", " R ", " A "}[e]
}

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

type Entity struct {
	Etype      EntityType
	Belongs_to Player
}

func (e Entity) String() string {
	return fmt.Sprintf("Type: %v, BelongsTo: %v", e.Etype, e.Belongs_to)
}

type Player struct {
	Name  string
	Color Color
}

func (p Player) String() string {
	return p.ColorString(p.Name)
}
