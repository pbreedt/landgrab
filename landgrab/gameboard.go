package landgrab

import (
	"fmt"
	"log"

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

	return gb
}

func (gb *Gameboard) CurrentPlayer() *Player {
	return &gb.Players[gb.currentPlayerIndex]
}

func (gb *Gameboard) NextPlayer() *Player {
	gb.currentPlayerIndex++
	if gb.currentPlayerIndex >= len(gb.Players) {
		gb.currentPlayerIndex = 0
	}
	return &gb.Players[gb.currentPlayerIndex]
}

func (gb *Gameboard) Display() *LandPiece {
	fmt.Printf("cur.piece:%d, cur.player.idx:%d, cur.player:%s\n", gb.currentPieceIndex, gb.currentPlayerIndex, gb.CurrentPlayer().Name)
	fmt.Println(gb)
	curPcIdx := gb.LandPieces.PrintUnplacedN(gb.currentPieceIndex, gb.showNumPieces)
	return &(*gb.LandPieces)[curPcIdx]
}

func (gb *Gameboard) String() string {
	bsh := ""
	for i, p := range gb.Players {
		bsh += fmt.Sprintf("Player %d: %v\n", i+1, p)
	}
	bs := ""
	for y := 0; y < boardsize; y++ {
		for x := 0; x < boardsize; x++ {
			bs += fmt.Sprintf("[%+v]", gb.Board[y][x])
		}
		bs += "\n"
	}
	return bsh + "\n" + bs
}

type Block struct {
	Marker         string
	Belongs_to     *Player
	HighlightErr   bool
	HighlightOK    bool
	LandPieceValue uint16
}

func (b Block) String() string {
	if b.HighlightErr {
		return ColorString(fmt.Sprintf("%1s", b.Marker), Red)
	}
	if b.HighlightOK {
		return ColorString(fmt.Sprintf("%1s", b.Marker), Green)
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
	Black  Color = iota // unused
	White               // default color
	Red                 // highlight: error
	Green               // highlight: ok
	Yellow              // player colors
	Blue
	Magenta
	Cyan
)

func (c Color) String() string {
	return [...]string{"Black", "White", "Red", "Green", "Yellow", "Blue", "Magenta", "Cyan"}[c]
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

func (gb *Gameboard) RemoveLandPieceFromGameboard(landPieceIdx int) {
	if (*gb.LandPieces)[landPieceIdx].PlacedAt == nil {
		return
	}

	lp := (*gb.LandPieces)[landPieceIdx]

	lpStr := lp.String()

	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			sidx := x + (y * 4)
			if lpStr[sidx:sidx+1] == "#" && (lp.PlacedAt.Y+y >= 0) && (lp.PlacedAt.X+x >= 0) && (lp.PlacedAt.Y+y < 16) && (lp.PlacedAt.X+x < 16) {
				gb.Board[lp.PlacedAt.Y+y][lp.PlacedAt.X+x] = Block{}
			}
		}
	}

	(*gb.LandPieces)[landPieceIdx].PlacedAt = nil
	// log.Default().Printf("Land piece:%s mem:%p, placed:%v\n", (*gb.LandPieces)[landPieceIdx], &(*gb.LandPieces)[landPieceIdx], (*gb.LandPieces)[landPieceIdx].PlacedAt)
}

// Move LandPiece from used index to just before current piece index
// then set current piece index 1 back
func (gb *Gameboard) SetNextLandPieceTo(landPieceIdx int) {
	if landPieceIdx < 0 || landPieceIdx >= len(*gb.LandPieces) {
		return
	}

	removedLandPiece := (*gb.LandPieces)[landPieceIdx]
	// Remove LandPiece from slice
	(*gb.LandPieces) = append((*gb.LandPieces)[:landPieceIdx], (*gb.LandPieces)[landPieceIdx+1:]...)

	// Add LandPiece back at index before current piece index (since slice is shorter now)
	(*gb.LandPieces) = append((*gb.LandPieces)[:gb.currentPieceIndex-1], append(LandPiecesSlice{removedLandPiece}, (*gb.LandPieces)[gb.currentPieceIndex-1:]...)...)

	gb.currentPieceIndex--
}

// MarkArea places provided marker in all blocks of the provided area
// For testing only
func (g *Gameboard) MarkArea(area Area, marker string) {
	b := Block{Marker: marker}

	log.Default().Printf("Marking area: %s with marker %s\n", area, marker)
	for x := area.Start.X; x <= area.End.X; x++ {
		for y := area.Start.Y; y <= area.End.Y; y++ {
			g.Board[y][x] = b
		}
	}
}
