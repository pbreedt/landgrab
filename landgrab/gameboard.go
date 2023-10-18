package landgrab

import (
	"fmt"

	"github.com/fatih/color"
)

const boardsize int = 12

type Gameboard struct {
	Board              [boardsize][boardsize]Block
	Players            []Player
	LandPieces         *LandPiecesSlice
	currentPieceIndex  int
	currentPlayerIndex int
	showNumPieces      int
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

func (gb Gameboard) CurrentPlayer() *Player {
	return &gb.Players[gb.currentPlayerIndex]
}

func (gb *Gameboard) NextPlayer() *Player {
	gb.currentPlayerIndex++
	if gb.currentPlayerIndex >= len(gb.Players) {
		gb.currentPlayerIndex = 0
	}
	return &gb.Players[gb.currentPlayerIndex]
}

func (gb Gameboard) Display() *LandPiece {
	fmt.Println(gb)
	curPcIdx := gb.LandPieces.PrintUnplacedN(gb.currentPieceIndex, gb.showNumPieces)
	return &(*gb.LandPieces)[curPcIdx]
}

func (gb Gameboard) String() string {
	bsh := ""
	for i, p := range gb.Players {
		bsh += fmt.Sprintf("Player %d: %v\n", i, p)
	}
	bs := ""
	for y := 0; y < boardsize; y++ {
		for x := 0; x < boardsize; x++ {
			// bs += fmt.Sprintf("[%+v](%2d,%2d)", gb.board[y][x], x, y)
			bs += fmt.Sprintf("[%+v]", gb.Board[y][x])
		}
		bs += "\n"
	}
	return bsh + "\n" + bs
}

type Block struct {
	Marker     string
	Belongs_to *Player
	Invalid    bool
}

func (b Block) String() string {
	if b.Invalid {
		return ColorString(fmt.Sprintf("%1s", b.Marker), Red)
	}
	if b.Belongs_to != nil {
		return ColorString(fmt.Sprintf("%1s", b.Marker), b.Belongs_to.Color)
	}
	return fmt.Sprintf("%1s", b.Marker)
}

var (
	OpenBlock      Block = Block{Marker: ""}
	SwapBlock      Block = Block{Marker: "S"}
	GrabBlock      Block = Block{Marker: "G"}
	RockBlock      Block = Block{Marker: "R"}
	HomeBlock      Block = Block{Marker: "H"}
	LandPieceBlock Block = Block{Marker: "#"}
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
	GrabCard
	RockCard
)

func (c Card) String() string {
	return [...]string{"Swap", "Grab", "Rock"}[c]
}

type Player struct {
	Name      string
	Color     Color
	SwapCards []Card
	GrabCards []Card
	RockCards []Card
}

func (p Player) String() string {
	return ColorString(fmt.Sprintf("%10s | Swap Cards:%d, Grab Cards:%d, Rock Cards:%d", p.Name, len(p.SwapCards), len(p.GrabCards), len(p.RockCards)), p.Color)
}
